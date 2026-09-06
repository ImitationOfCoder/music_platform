CREATE TABLE IF NOT EXISTS music_platform.profile_avatars (
    id BIGSERIAL PRIMARY KEY,
    profile_id BIGINT NOT NULL,

    filename VARCHAR(255) UNIQUE NOT NULL,
    size INTEGER NOT NULL,
    color VARCHAR(255) NOT NULL,

    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_profiles_avatar FOREIGN KEY (profile_id) REFERENCES music_platform.profiles(id) ON DELETE CASCADE ON UPDATE CASCADE
);
