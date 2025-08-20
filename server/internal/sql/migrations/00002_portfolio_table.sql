-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS portfolio(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id NOT NULL REFERENCES users(id) ON DELETE CASCADE, 
    name TEXT NOT NULL, 
    invested_value TEXT, 
    current_value TEXT ,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS portfolio;
-- +goose StatementEnd
