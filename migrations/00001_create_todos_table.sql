-- +goose Up
CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    completed BOOLEAN DEFAULT FALSE,
    createdAt TIMESTAMP DEFAULT NOW(),
    completedAt TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS todos;

