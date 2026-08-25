package security

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type Argon2Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

func HashPassword(password string) (string, error) {
	params := &Argon2Params{
		Memory:      64 * 1024, // 64 MB
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	// Генерируем соль
	salt := make([]byte, params.SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("failed to generate salt: %w", err)
	}

	// Хешируем пароль
	hash := argon2.IDKey(
		[]byte(password),
		salt,
		params.Iterations,
		params.Memory,
		params.Parallelism,
		params.KeyLength,
	)

	// Кодируем в строку формата:
	// $argon2id$v=19$m=65536,t=3,p=2$c29tZXNhbHQ$hash
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encoded := fmt.Sprintf(
		"$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		params.Memory,
		params.Iterations,
		params.Parallelism,
		b64Salt,
		b64Hash,
	)

	return encoded, nil
}

func VerifyPassword(password string, encodedHash string) (bool, error) {
	// Парсим параметры из строки
	params, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	// Вычисляем хеш для переданного пароля
	otherHash := argon2.IDKey(
		[]byte(password),
		salt,
		params.Iterations,
		params.Memory,
		params.Parallelism,
		params.KeyLength,
	)

	// Сравниваем хеши (constant-time сравнение)
	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}

	return false, nil
}

func decodeHash(encodedHash string) (*Argon2Params, []byte, []byte, error) {
	// Формат: $argon2id$v=19$m=65536,t=3,p=2$salt$hash
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return nil, nil, nil, errors.New("invalid hash format")
	}

	if parts[1] != "argon2id" {
		return nil, nil, nil, errors.New("invalid algorithm (expected argon2id)")
	}

	// Версия
	version := strings.Split(parts[2], "=")
	if len(version) != 2 || version[1] != "19" {
		return nil, nil, nil, errors.New("invalid version (expected v=19)")
	}

	// Параметры: m=65536,t=3,p=2
	paramsMap := make(map[string]uint32)
	paramParts := strings.Split(parts[3], ",")
	for _, p := range paramParts {
		kv := strings.Split(p, "=")
		if len(kv) != 2 {
			return nil, nil, nil, errors.New("invalid parameter format")
		}
		var value uint32
		if _, err := fmt.Sscanf(kv[1], "%d", &value); err != nil {
			return nil, nil, nil, fmt.Errorf("invalid parameter value: %w", err)
		}
		paramsMap[kv[0]] = value
	}

	params := &Argon2Params{
		Memory:      paramsMap["m"],
		Iterations:  paramsMap["t"],
		Parallelism: uint8(paramsMap["p"]),
	}

	// Декодируем соль
	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to decode salt: %w", err)
	}

	// Декодируем хеш
	hash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to decode hash: %w", err)
	}

	params.SaltLength = uint32(len(salt))
	params.KeyLength = uint32(len(hash))

	return params, salt, hash, nil
}
