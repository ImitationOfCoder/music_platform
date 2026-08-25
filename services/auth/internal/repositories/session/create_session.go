package session_repository

import (
	"context"
	"fmt"
)

func (r *SessionRepository) CreateSession(ctx context.Context, userId int64) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		INSERT INTO music_platform.sessions (user_id)
		VALUES ($1)
		RETURNING token;
	`

	var token string
	err := r.pool.QueryRow(
		ctx,
		query,
		userId,
	).Scan(&token)
	if err != nil {
		return "", fmt.Errorf("SessionRepository - CreateSession - Scan: %w", err)
	}

	return token, nil
}
