-- name: CreatePortfolio :one
INSERT INTO portfolio (
    id,user_id,name,invested_value,current_value
) VALUES ($1,$2,$3,$4,$5) RETURNING id,user_id,name,invested_value,current_value,created_at,updated_at;