package security

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// Argon2Config содержит параметры для хеширования
type Argon2Config struct {
	Time    uint32 // количество итераций
	Memory  uint32 // память в килобайтах
	Threads uint8  // количество потоков
	KeyLen  uint32 // длина выходного ключа
	SaltLen uint32 // длина соли
}

// DefaultConfig возвращает конфигурацию по умолчанию
func DefaultConfig() *Argon2Config {
	return &Argon2Config{
		Time:    1,
		Memory:  64 * 1024, // 64 MB
		Threads: 4,
		KeyLen:  32,
		SaltLen: 16,
	}
}

// PasswordHasher реализует хеширование паролей через Argon2id
type PasswordHasher struct {
	config *Argon2Config
}

// NewPasswordHasher создает новый экземпляр PasswordHasher
func NewPasswordHasher(config *Argon2Config) *PasswordHasher {
	if config == nil {
		config = DefaultConfig()
	}
	return &PasswordHasher{config: config}
}

// generateSalt генерирует случайную соль
func (h *PasswordHasher) generateSalt() ([]byte, error) {
	salt := make([]byte, h.config.SaltLen)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, fmt.Errorf("не удалось сгенерировать соль: %w", err)
	}
	return salt, nil
}

// HashPassword хеширует пароль с использованием Argon2id
func (h *PasswordHasher) HashPassword(password string) (string, error) {
	// Генерируем случайную соль
	salt, err := h.generateSalt()
	if err != nil {
		return "", err
	}

	// Хешируем пароль
	hash := argon2.IDKey(
		[]byte(password),
		salt,
		h.config.Time,
		h.config.Memory,
		h.config.Threads,
		h.config.KeyLen,
	)

	// Кодируем в формат для хранения
	// Формат: $argon2id$v=19$m=65536,t=1,p=4$<salt>$<hash>
	saltEncoded := base64.RawStdEncoding.EncodeToString(salt)
	hashEncoded := base64.RawStdEncoding.EncodeToString(hash)

	return fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		h.config.Memory,
		h.config.Time,
		h.config.Threads,
		saltEncoded,
		hashEncoded,
	), nil
}

// VerifyPassword проверяет пароль против хеша
func (h *PasswordHasher) VerifyPassword(password, hash string) (bool, error) {
	// Парсим строку хеша
	parts := strings.Split(hash, "$")
	if len(parts) != 6 {
		return false, fmt.Errorf("неверный формат хеша")
	}

	// Проверяем версию
	if parts[1] != "argon2id" {
		return false, fmt.Errorf("неподдерживаемый алгоритм")
	}

	// Парсим параметры
	versionPart := strings.Split(parts[2], "=")
	if len(versionPart) != 2 || versionPart[0] != "v" {
		return false, fmt.Errorf("неверный формат версии")
	}

	paramsPart := strings.Split(parts[3], ",")
	var memory, time, threads uint32
	for _, param := range paramsPart {
		kv := strings.Split(param, "=")
		if len(kv) != 2 {
			continue
		}
		switch kv[0] {
		case "m":
			fmt.Sscanf(kv[1], "%d", &memory)
		case "t":
			fmt.Sscanf(kv[1], "%d", &time)
		case "p":
			fmt.Sscanf(kv[1], "%d", &threads)
		}
	}

	// Декодируем соль и хеш
	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, fmt.Errorf("не удалось декодировать соль: %w", err)
	}

	storedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, fmt.Errorf("не удалось декодировать хеш: %w", err)
	}

	// Вычисляем хеш для проверки
	computedHash := argon2.IDKey(
		[]byte(password),
		salt,
		time,
		memory,
		uint8(threads),
		uint32(len(storedHash)),
	)

	// Сравниваем хеши
	if subtle.ConstantTimeCompare(computedHash, storedHash) == 1 {
		return true, nil
	}

	return false, nil
}

// VerifyPasswordHash удобная функция для проверки пароля
func VerifyPasswordHash(password, hash string) (bool, error) {
	hasher := NewPasswordHasher(nil)
	return hasher.VerifyPassword(password, hash)
}

// HashPassword удобная функция для хеширования пароля
func HashPassword(password string) (string, error) {
	hasher := NewPasswordHasher(nil)
	return hasher.HashPassword(password)
}
