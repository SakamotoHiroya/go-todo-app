-- ============================
-- Table: users
-- ============================

-- name: CreateUser :one
INSERT INTO users (
  provider,
  provider_user_id,
  email,
  name,
  status
) VALUES (
  $1,  -- provider (oauth_provider)
  $2,  -- provider_user_id
  $3,  -- email
  $4,  -- name
  $5   -- status (user_status)
)
RETURNING
  id,
  provider,
  provider_user_id,
  email,
  name,
  status,
  created_at,
  updated_at;

-- name: GetUserByID :one
SELECT
  id,
  provider,
  provider_user_id,
  email,
  name,
  status,
  created_at,
  updated_at
FROM users
WHERE id = $1;  -- id

-- name: GetUserByEmail :one
SELECT
  id,
  provider,
  provider_user_id,
  email,
  name,
  status,
  created_at,
  updated_at
FROM users
WHERE email = $1;  -- email

-- name: ListUsers :many
SELECT
  id,
  provider,
  provider_user_id,
  email,
  name,
  status,
  created_at,
  updated_at
FROM users
ORDER BY created_at DESC;

-- name: UpdateUser :one
UPDATE users
SET
  provider          = $2,  -- provider
  provider_user_id  = $3,  -- provider_user_id
  email             = $4,  -- email
  name              = $5,  -- name
  status            = $6,  -- status (user_status)
  updated_at        = NOW()
WHERE id = $1        -- id
RETURNING
  id,
  provider,
  provider_user_id,
  email,
  name,
  status,
  created_at,
  updated_at;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;  -- id