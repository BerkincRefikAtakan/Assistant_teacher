-- name: GetTeachers :many
SELECT id, name, surname, created_at
FROM teachers;

-- name: GetTeacherByID :one
SELECT id, name, surname, created_at
FROM teachers
WHERE id = $1;

-- name: CreateTeacher :one
INSERT INTO teachers (name, surname)
VALUES ($1, $2)
RETURNING id, name, surname, created_at;

-- name: UpdateTeacher :one
UPDATE teachers
SET name = $1, surname = $2
WHERE id = $3
RETURNING id, name, surname, created_at;

-- name: DeleteTeacher :exec
DELETE FROM teachers
WHERE id = $1;
