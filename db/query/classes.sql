-- name: GetClasses :many
SELECT id, teacher_id, class_name, created_at
FROM classes;

-- name: GetClassByID :one
SELECT id, teacher_id, class_name, created_at
FROM classes
WHERE id = $1;

-- name: GetClassesByTeacher :many
SELECT id, teacher_id, class_name, created_at
FROM classes
WHERE teacher_id = $1;

-- name: CreateClass :one
INSERT INTO classes (teacher_id, class_name)
VALUES ($1, $2)
RETURNING id, teacher_id, class_name, created_at;

-- name: UpdateClass :one
UPDATE classes
SET teacher_id = $1, class_name = $2
WHERE id = $3
RETURNING id, teacher_id, class_name, created_at;

-- name: DeleteClass :exec
DELETE FROM classes
WHERE id = $1;
