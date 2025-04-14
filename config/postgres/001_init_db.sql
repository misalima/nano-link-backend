-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS urls (
                      id SERIAL PRIMARY KEY,
                      short_id VARCHAR(10) UNIQUE,
                      original_url TEXT NOT NULL,
                      visits INT DEFAULT 0,
                      created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS users (
                      id SERIAL PRIMARY KEY,
                      username VARCHAR(50) UNIQUE NOT NULL,
                      password_hash TEXT NOT NULL,
                      created_at TIMESTAMP DEFAULT NOW()
);
---- create above / drop below ----

    DROP TABLE IF EXISTS urls;
    DROP TABLE IF EXISTS users;
-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
