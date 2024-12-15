-- name: GetClasses :one
SELECT id, teacher_id, class_name, created_at
FROM classes
WHERE class_name = $1 AND teacher_id = $2;

-- name: CreateClass :one
INSERT INTO classes (teacher_id, class_name)
VALUES ($1, $2)
RETURNING id, teacher_id, class_name, created_at;

-- name: UpdateClass :one
UPDATE classes
SET teacher_id = $1
WHERE id = $2
RETURNING id, teacher_id, class_name, created_at;

-- name: DeleteClass :exec
DELETE FROM classes
WHERE id = $1;
