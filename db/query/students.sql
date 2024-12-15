-- name: GetStudents :one
SELECT id, class_id, name, surname, created_at
FROM students
WHERE name = $1 AND surname = $2 AND class_id = $3;

-- name: GetStudentsByClass :many
SELECT id, class_id, name, surname, created_at
FROM students
WHERE class_id = $1
LIMIT $2 OFFSET $3;

-- name: CreateStudent :one
INSERT INTO students (class_id, name, surname)
VALUES ($1, $2, $3)
RETURNING id, class_id, name, surname, created_at;

-- name: UpdateStudent :one
UPDATE students
SET class_id = $1, name = $2, surname = $3
WHERE id = $4
RETURNING id, class_id, name, surname, created_at;

-- name: DeleteStudent :exec
DELETE FROM students
WHERE id = $1;

