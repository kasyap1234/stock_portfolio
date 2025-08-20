-- name: CreateUser :one 
INSERT INTO users (
    id, name, username, password 
) VALUES ($1,$2,$3,$4)
RETURNING id,name,username, password,created_at; 

-- name: FindUserByUsername :one 
SELECT id,name,username,password,created_at FROM users WHERE username=$1; 
