CREATE TABLE IF NOT EXISTS music_platform.profiles (
    id BIGSERIAL PRIMARY KEY,
    account_id BIGINT NOT NULL REFERENCES music_platform.accounts(id) ON DELETE CASCADE ON UPDATE CASCADE,
    name VARCHAR(32) NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT valid_name_check CHECK (char_length(name) BETWEEN 1 AND 32)
);

CREATE INDEX idx_profiles_account_id ON music_platform.profiles(account_id);
