DROP TABLE IF EXISTS blogs CASCADE;
DROP TABLE IF EXISTS news CASCADE;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS CITEXT;

-- Create Blogs Table with Timestamps
CREATE TABLE blogs (
                       id SERIAL PRIMARY KEY,
                       title VARCHAR(255) NOT NULL,
                       content TEXT NOT NULL,
                       created_at TIMESTAMPTZ DEFAULT current_timestamp NOT NULL,
                       updated_at TIMESTAMPTZ,
                       deleted_at TIMESTAMPTZ
);
CREATE INDEX idx_blogs_title ON blogs(title);

-- Create News Table with Timestamps
CREATE TABLE news (
                      id SERIAL PRIMARY KEY,
                      title VARCHAR(255) NOT NULL,
                      content TEXT NOT NULL,
                      created_at TIMESTAMPTZ DEFAULT current_timestamp NOT NULL,
                      updated_at TIMESTAMPTZ,
                      deleted_at TIMESTAMPTZ
);
CREATE INDEX idx_news_title ON news(title);
