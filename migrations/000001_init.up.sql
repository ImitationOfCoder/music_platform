CREATE SCHEMA IF NOT EXISTS music_platform;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS music_platform.users (
	id BIGSERIAL PRIMARY KEY,
	name VARCHAR(32) NOT NULL,
	email VARCHAR(254) NOT NULL UNIQUE,
	password_hash VARCHAR(255) NOT NULL,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

	CONSTRAINT valid_name_check CHECK (char_length(name) BETWEEN 1 AND 32),
	CONSTRAINT valid_email_check CHECK (email ~ '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$')
);

CREATE INDEX idx_users_id ON music_platform.users(id);
CREATE INDEX idx_users_email ON music_platform.users(email);

CREATE TABLE IF NOT EXISTS music_platform.sessions (
	id BIGSERIAL PRIMARY KEY,
	user_id BIGINT NOT NULL,
	token UUID DEFAULT uuid_generate_v4() NOT NULL UNIQUE,

	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

 	CONSTRAINT fk_sessions_user FOREIGN KEY (user_id) REFERENCES music_platform.users(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE INDEX idx_sessions_token ON music_platform.sessions(token);
CREATE INDEX idx_sessions_user_id ON music_platform.sessions(user_id);
