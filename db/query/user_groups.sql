-- name: ListUserGroups :many
SELECT
  id,
  name
FROM user_groups
ORDER BY id;

-- name: GetUserGroupByID :one
SELECT
  id,
  name
FROM user_groups
WHERE id = $1;

-- name: CreateUserGroup :one
INSERT INTO user_groups (
  name
) VALUES (
  $1
)
RETURNING
  id,
  name;

-- name: UpdateUserGroup :exec
UPDATE user_groups
SET
  name = $1
WHERE
  id = $2;

-- name: DeleteUserGroup :exec
DELETE FROM user_groups
WHERE
  id = $1;

-- name: AddUserToGroup :exec
INSERT INTO user_group_members (
  user_id,
  group_id
) VALUES (
  $1, $2
);

-- name: RemoveUserFromGroup :exec
DELETE FROM user_group_members
WHERE
  user_id  = $1
  AND group_id = $2;

-- name: ListGroupsByUser :many
SELECT
  g.id,
  g.name
FROM user_groups AS g
JOIN user_group_members AS m
  ON g.id = m.group_id
WHERE
  m.user_id = $1
ORDER BY g.id;

-- name: ListUsersByGroup :many
SELECT
  u.id,
  u.provider,
  u.provider_user_id,
  u.email,
  u.name,
  u.avatar_url,
  u.status,
  u.created_at,
  u.updated_at
FROM users AS u
JOIN user_group_members AS m
  ON u.id = m.user_id
WHERE
  m.group_id = $1
ORDER BY u.id;