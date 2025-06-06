// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: task_groups_queries.sql

package db

import (
	"context"
	"database/sql"
)

const addGroupMember = `-- name: AddGroupMember :one


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
  user_id
`

type AddGroupMemberParams struct {
	TaskGroupID int64 `json:"task_group_id"`
	UserID      int64 `json:"user_id"`
}

// id
// ============================
// Table: task_group_members
// ============================
func (q *Queries) AddGroupMember(ctx context.Context, arg AddGroupMemberParams) (TaskGroupMember, error) {
	row := q.db.QueryRowContext(ctx, addGroupMember, arg.TaskGroupID, arg.UserID)
	var i TaskGroupMember
	err := row.Scan(&i.ID, &i.TaskGroupID, &i.UserID)
	return i, err
}

const createTaskGroup = `-- name: CreateTaskGroup :one

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
  owner_user_id
`

type CreateTaskGroupParams struct {
	Name        string        `json:"name"`
	OwnerUserID sql.NullInt64 `json:"owner_user_id"`
}

// ============================
// Table: task_groups
// ============================
func (q *Queries) CreateTaskGroup(ctx context.Context, arg CreateTaskGroupParams) (TaskGroup, error) {
	row := q.db.QueryRowContext(ctx, createTaskGroup, arg.Name, arg.OwnerUserID)
	var i TaskGroup
	err := row.Scan(&i.ID, &i.Name, &i.OwnerUserID)
	return i, err
}

const deleteTaskGroup = `-- name: DeleteTaskGroup :exec
DELETE FROM task_groups
WHERE id = $1
`

func (q *Queries) DeleteTaskGroup(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTaskGroup, id)
	return err
}

const getMembersByGroupID = `-- name: GetMembersByGroupID :many
SELECT
  id,
  task_group_id,
  user_id
FROM task_group_members
WHERE task_group_id = $1
`

func (q *Queries) GetMembersByGroupID(ctx context.Context, taskGroupID int64) ([]TaskGroupMember, error) {
	rows, err := q.db.QueryContext(ctx, getMembersByGroupID, taskGroupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TaskGroupMember{}
	for rows.Next() {
		var i TaskGroupMember
		if err := rows.Scan(&i.ID, &i.TaskGroupID, &i.UserID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTaskGroupByID = `-- name: GetTaskGroupByID :one
SELECT
  id,
  name,
  owner_user_id
FROM task_groups
WHERE id = $1
`

func (q *Queries) GetTaskGroupByID(ctx context.Context, id int64) (TaskGroup, error) {
	row := q.db.QueryRowContext(ctx, getTaskGroupByID, id)
	var i TaskGroup
	err := row.Scan(&i.ID, &i.Name, &i.OwnerUserID)
	return i, err
}

const listGroupMembers = `-- name: ListGroupMembers :many

SELECT
  id,
  task_group_id,
  user_id
FROM task_group_members
ORDER BY id
`

// task_group_id
func (q *Queries) ListGroupMembers(ctx context.Context) ([]TaskGroupMember, error) {
	rows, err := q.db.QueryContext(ctx, listGroupMembers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TaskGroupMember{}
	for rows.Next() {
		var i TaskGroupMember
		if err := rows.Scan(&i.ID, &i.TaskGroupID, &i.UserID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTaskGroups = `-- name: ListTaskGroups :many

SELECT
  id,
  name,
  owner_user_id
FROM task_groups
ORDER BY name
`

// id
func (q *Queries) ListTaskGroups(ctx context.Context) ([]TaskGroup, error) {
	rows, err := q.db.QueryContext(ctx, listTaskGroups)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TaskGroup{}
	for rows.Next() {
		var i TaskGroup
		if err := rows.Scan(&i.ID, &i.Name, &i.OwnerUserID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeGroupMember = `-- name: RemoveGroupMember :exec
DELETE FROM task_group_members
WHERE id = $1
`

func (q *Queries) RemoveGroupMember(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, removeGroupMember, id)
	return err
}

const updateTaskGroup = `-- name: UpdateTaskGroup :one
UPDATE task_groups
SET
  name           = $2,  -- name
  owner_user_id  = $3   -- owner_user_id
WHERE id = $1        -- id
RETURNING
  id,
  name,
  owner_user_id
`

type UpdateTaskGroupParams struct {
	ID          int64         `json:"id"`
	Name        string        `json:"name"`
	OwnerUserID sql.NullInt64 `json:"owner_user_id"`
}

func (q *Queries) UpdateTaskGroup(ctx context.Context, arg UpdateTaskGroupParams) (TaskGroup, error) {
	row := q.db.QueryRowContext(ctx, updateTaskGroup, arg.ID, arg.Name, arg.OwnerUserID)
	var i TaskGroup
	err := row.Scan(&i.ID, &i.Name, &i.OwnerUserID)
	return i, err
}
