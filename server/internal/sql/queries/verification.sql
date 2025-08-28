-- name: CreateVerificationToken :one
INSERT INTO verification_tokens (
    user_id, token, expires_at
) VALUES (
    $1, $2, $3
) RETURNING id, user_id, token, expires_at, created_at;

-- name: GetVerificationToken :one
SELECT id, user_id, token, expires_at, created_at
FROM verification_tokens
WHERE token = $1;

-- name: DeleteVerificationToken :exec
DELETE FROM verification_tokens
WHERE id = $1;
