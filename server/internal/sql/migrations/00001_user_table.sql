-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
id  UUID PRIMARY KEY DEFAULT   gen_random_uuid(),
name TEXT NOT NULL, 
email TEXT NOT NULL,
password TEXT NOT NULL, 
created_at TIMESTAMPTZ DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users; 
-- +goose StatementEnd
