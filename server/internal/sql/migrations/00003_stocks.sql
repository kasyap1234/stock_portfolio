-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS stocks(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    portfolio_id UUID NOT NULL REFERENCES portfolio(id) ON DELETE CASCADE, 
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    symbol TEXT NOT NULL, 
    name TEXT NOT NULL, 
    quantity NUMERIC  NOT NULL, 
    purchase_price NUMERIC NOT NULL,
    purchase_date DATE NOT NULL, 
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS stocks; 
-- +goose StatementEnd
