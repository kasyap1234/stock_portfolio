-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
 id  SERIAL PRIMARY KEY,
name TEXT NOT NULL, 
username TEXT NOT NULL, 
password TEXT NOT NULL, 
created_at TIMESTAMPTZ DEFAULT now()
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users; 
-- +goose StatementEnd
