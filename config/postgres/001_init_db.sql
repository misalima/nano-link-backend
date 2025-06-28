-- Write your migrate up statements here
-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Users Table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
    );

-- URLs Table
CREATE TABLE IF NOT EXISTS urls (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    short_id VARCHAR(15) UNIQUE NOT NULL,
    custom_short_id VARCHAR(15) UNIQUE,
    original_url VARCHAR(2048) NOT NULL,
    total_visits INT DEFAULT 0 CHECK (total_visits >= 0),
    user_id UUID,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    );

-- Tags Table
CREATE TABLE IF NOT EXISTS tags (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(50) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
    );

-- URL Tags Table (Join Table)
CREATE TABLE IF NOT EXISTS url_tags (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    url_id UUID NOT NULL,
    tag_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (url_id) REFERENCES urls(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
    );

-- URL Visits Table
CREATE TABLE IF NOT EXISTS url_visits (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    url_id UUID NOT NULL,
    visited_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (url_id) REFERENCES urls(id) ON DELETE CASCADE
    );

-- Function to increment total_visits
CREATE OR REPLACE FUNCTION increment_total_visits()
RETURNS TRIGGER AS $$
BEGIN
UPDATE urls
SET total_visits = total_visits + 1
WHERE id = NEW.url_id;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger to call the function on insert into url_visits
CREATE TRIGGER trigger_increment_total_visits
    AFTER INSERT ON url_visits
    FOR EACH ROW
    EXECUTE FUNCTION increment_total_visits();

-- Write your migrate down statements here
-- Drop the trigger and function
DROP TRIGGER IF EXISTS trigger_increment_total_visits ON url_visits;
DROP FUNCTION IF EXISTS increment_total_visits;

-- Drop tables in reverse order of creation
DROP TABLE IF EXISTS url_visits;
DROP TABLE IF EXISTS url_tags;
DROP TABLE IF EXISTS tags;
DROP TABLE IF EXISTS urls;
DROP TABLE IF EXISTS users;

-- Drop the UUID extension
DROP EXTENSION IF EXISTS "uuid-ossp";