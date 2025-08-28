-- name: CreateUser :one 
INSERT INTO users (
    id, name, email, password
) VALUES ($1,$2,$3,$4)
RETURNING id,name,email, password,created_at,email_verified; 

-- name: GetUserByEmail :one
SELECT id,name,email,password,created_at,email_verified FROM users WHERE email=$1;

-- name: GetUserByID :one
SELECT id,name,email,password,created_at,email_verified FROM users WHERE id=$1;

-- name: VerifyUserEmail :exec
UPDATE users SET email_verified = true WHERE email = $1;
