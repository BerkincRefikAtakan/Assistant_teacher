// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: classes.sql

package db

import (
	"context"
)

const createClass = `-- name: CreateClass :one
INSERT INTO classes (teacher_id, class_name)
VALUES ($1, $2)
RETURNING id, teacher_id, class_name, created_at
`

type CreateClassParams struct {
	TeacherID int64  `json:"teacher_id"`
	ClassName string `json:"class_name"`
}

func (q *Queries) CreateClass(ctx context.Context, arg CreateClassParams) (Class, error) {
	row := q.db.QueryRow(ctx, createClass, arg.TeacherID, arg.ClassName)
	var i Class
	err := row.Scan(
		&i.ID,
		&i.TeacherID,
		&i.ClassName,
		&i.CreatedAt,
	)
	return i, err
}

const deleteClass = `-- name: DeleteClass :exec
DELETE FROM classes
WHERE id = $1
`

func (q *Queries) DeleteClass(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteClass, id)
	return err
}

const getClasses = `-- name: GetClasses :one
SELECT id, teacher_id, class_name, created_at
FROM classes
WHERE class_name = $1 AND teacher_id = $2
`

type GetClassesParams struct {
	ClassName string `json:"class_name"`
	TeacherID int64  `json:"teacher_id"`
}

func (q *Queries) GetClasses(ctx context.Context, arg GetClassesParams) (Class, error) {
	row := q.db.QueryRow(ctx, getClasses, arg.ClassName, arg.TeacherID)
	var i Class
	err := row.Scan(
		&i.ID,
		&i.TeacherID,
		&i.ClassName,
		&i.CreatedAt,
	)
	return i, err
}

const updateClass = `-- name: UpdateClass :one
UPDATE classes
SET teacher_id = $1
WHERE id = $2
RETURNING id, teacher_id, class_name, created_at
`

type UpdateClassParams struct {
	TeacherID int64 `json:"teacher_id"`
	ID        int64 `json:"id"`
}

func (q *Queries) UpdateClass(ctx context.Context, arg UpdateClassParams) (Class, error) {
	row := q.db.QueryRow(ctx, updateClass, arg.TeacherID, arg.ID)
	var i Class
	err := row.Scan(
		&i.ID,
		&i.TeacherID,
		&i.ClassName,
		&i.CreatedAt,
	)
	return i, err
}
