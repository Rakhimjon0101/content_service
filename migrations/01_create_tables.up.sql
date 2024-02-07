DROP TABLE IF EXISTS blogs CASCADE;
DROP TABLE IF EXISTS news CASCADE;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create Blogs Table with Timestamps
CREATE TABLE blogs (
                       id uuid not null constraint blogs_pk primary key DEFAULT (uuid_generate_v4()),
                       title VARCHAR(255) NOT NULL,
                       content TEXT NOT NULL,
                       created_at TIMESTAMPTZ DEFAULT current_timestamp NOT NULL,
                       updated_at TIMESTAMPTZ,
                       deleted_at TIMESTAMPTZ
);
CREATE INDEX idx_blogs_title ON blogs(title);

-- Create News Table with Timestamps
CREATE TABLE news (
                      id uuid not null constraint news_pk primary key DEFAULT (uuid_generate_v4()),
                      title VARCHAR(255) NOT NULL,
                      content TEXT NOT NULL,
                      created_at TIMESTAMPTZ DEFAULT current_timestamp NOT NULL,
                      updated_at TIMESTAMPTZ,
                      deleted_at TIMESTAMPTZ
);
CREATE INDEX idx_news_title ON news(title);
