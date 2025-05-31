-- name: ListUsers :many
SELECT
  id,
  provider,
  provider_user_id,
  email,
  name,
  avatar_url,
  status,
  created_at,
  updated_at
FROM users
ORDER BY id;

-- name: GetUserByID :one
SELECT
  id,
  provider,
  provider_user_id,
  email,
  name,
  avatar_url,
  status,
  created_at,
  updated_at
FROM users
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (
  provider,
  provider_user_id,
  email,
  name,
  avatar_url
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING
  id,
  provider,
  provider_user_id,
  email,
  name,
  avatar_url,
  status,
  created_at,
  updated_at;

-- name: LogicalDeleteUser :exec
UPDATE users
SET
  status = 'inactive'
WHERE
  id = $1;

-- name: UpdateUser :exec
UPDATE users
SET
  email      = $1,
  name       = $2,
  avatar_url = $3,
  updated_at = NOW()
WHERE
  id = $4;