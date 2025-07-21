-- +goose Up
CREATE TABLE posts(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT DEFAULT NULL,
    url TEXT UNIQUE NOT NULL,
    description TEXT DEFAULT NULL,
    published_at TIMESTAMP DEFAULT NULL,
    feed_id UUID NOT NULL
);

-- +goose Down
DROP TABLE posts;