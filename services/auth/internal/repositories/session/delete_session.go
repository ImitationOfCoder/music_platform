package session_repository

import (
	"context"
	"fmt"
)

func (r *SessionRepository) DeleteSession(ctx context.Context, token string) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		DELETE FROM music_platform.sessions
		WHERE token = $1
	`
	_, err := r.pool.Exec(ctx, query, token)
	if err != nil {
		return fmt.Errorf("SessionRepository - DeleteSession - Query: %w", err)
	}

	return nil
}
