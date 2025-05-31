-- ============================
-- Table: task_groups
-- ============================

-- name: CreateTaskGroup :one
INSERT INTO task_groups (
  name,
  owner_user_id
) VALUES (
  $1,  -- name
  $2   -- owner_user_id
)
RETURNING
  id,
  name,
  owner_user_id;

-- name: GetTaskGroupByID :one
SELECT
  id,
  name,
  owner_user_id
FROM task_groups
WHERE id = $1;  -- id

-- name: ListTaskGroups :many
SELECT
  id,
  name,
  owner_user_id
FROM task_groups
ORDER BY name;

-- name: UpdateTaskGroup :one
UPDATE task_groups
SET
  name           = $2,  -- name
  owner_user_id  = $3   -- owner_user_id
WHERE id = $1        -- id
RETURNING
  id,
  name,
  owner_user_id;

-- name: DeleteTaskGroup :exec
DELETE FROM task_groups
WHERE id = $1;  -- id

-- ============================
-- Table: task_group_members
-- ============================

-- name: AddGroupMember :one
INSERT INTO task_group_members (
  task_group_id,
  user_id
) VALUES (
  $1,  -- task_group_id
  $2   -- user_id
)
RETURNING
  id,
  task_group_id,
  user_id;

-- name: GetMembersByGroupID :many
SELECT
  id,
  task_group_id,
  user_id
FROM task_group_members
WHERE task_group_id = $1;  -- task_group_id

-- name: ListGroupMembers :many
SELECT
  id,
  task_group_id,
  user_id
FROM task_group_members
ORDER BY id;

-- name: RemoveGroupMember :exec
DELETE FROM task_group_members
WHERE id = $1;  -- id