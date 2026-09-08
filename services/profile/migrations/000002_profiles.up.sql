CREATE TABLE IF NOT EXISTS music_platform.profiles (
    id BIGINT PRIMARY KEY,
    account_id BIGINT NOT NULL,
    name VARCHAR(32) NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT valid_name_check CHECK (char_length(name) BETWEEN 1 AND 32)
);
