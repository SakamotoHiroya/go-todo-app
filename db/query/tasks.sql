-- name: ListTasks :many
SELECT
  id,
  title,
  detail,
  status,
  task_group_id,
  created_at,
  updated_at
FROM tasks
ORDER BY id;

-- name: GetTaskByID :one
SELECT
  id,
  title,
  detail,
  status,
  task_group_id,
  created_at,
  updated_at
FROM tasks
WHERE id = $1;

-- name: CreateTask :one
INSERT INTO tasks (
  title,
  detail,
  status,
  task_group_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING
  id,
  title,
  detail,
  status,
  task_group_id,
  created_at,
  updated_at;

-- name: UpdateTask :exec
UPDATE tasks
SET
  title          = $1,
  detail         = $2,
  status         = $3,
  task_group_id  = $4,
  updated_at     = NOW()
WHERE
  id = $5;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE
  id = $1;

-- name: AssignTaskToUser :exec
INSERT INTO task_assignees (
  task_id,
  user_id
) VALUES (
  $1, $2
);

-- name: UnassignTaskFromUser :exec
DELETE FROM task_assignees
WHERE
  task_id = $1
  AND user_id = $2;

-- name: ListUsersByTask :many
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
JOIN task_assignees AS a
  ON u.id = a.user_id
WHERE
  a.task_id = $1
ORDER BY u.id;

-- name: ListTasksByUser :many
SELECT
  t.id,
  t.title,
  t.detail,
  t.status,
  t.task_group_id,
  t.created_at,
  t.updated_at
FROM tasks AS t
JOIN task_assignees AS a
  ON t.id = a.task_id
WHERE
  a.user_id = $1
ORDER BY t.id;