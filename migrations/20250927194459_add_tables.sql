-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS notes (
    id INTEGER PRIMARY KEY,
    content TEXT NOT NULL,
    ttl INTEGER NOT NULL DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS notes;
-- +goose StatementEnd
