-- name: CreatePortfolio :one
INSERT INTO portfolio (
user_id,name,invested_value,current_value
) VALUES ($1,$2,$3,$4) RETURNING id,user_id,name,invested_value,current_value,created_at,updated_at;

-- name: ListPortfolio :many 
SELECT id,user_id,name,invested_value,current_value FROM portfolio where user_id=$1; 
