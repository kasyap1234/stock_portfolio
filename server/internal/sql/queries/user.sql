-- name: CreateUser :one 
INSERT INTO users (
    id, name, email, password 
) VALUES ($1,$2,$3,$4)
RETURNING id,name,email, password,created_at; 

-- name: GetUserByEmail :one
SELECT id,name,email,password,created_at FROM users WHERE email=$1;



