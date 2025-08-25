-- name: AddStock :one 
INSERT INTO stocks (
    portfolio_id,user_id,symbol,name,quantity,purchase_price,purchase_date) VALUES ($1,$2,$3,$4,$5,$6,$7)
RETURNING id,portfolio_id,user_id,symbol,name,quantity,purchase_price,purchase_date,created_at,updated_at;


-- name: DeleteStock :exec 
DELETE FROM stocks where id=$1 AND user_id=$2 AND portfolio_id=$3;

-- name: UpdateStock :one
UPDATE stocks
SET
    symbol = COALESCE($4, symbol),
    name = COALESCE($5, name),
    quantity = COALESCE($6, quantity),
    purchase_price = COALESCE($7, purchase_price),

    updated_at = now()
WHERE id = $1 AND portfolio_id = $2 AND user_id = $3
RETURNING id, portfolio_id, user_id, symbol, name, quantity, purchase_price, created_at, updated_at;
