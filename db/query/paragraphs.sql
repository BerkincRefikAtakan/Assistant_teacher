-- name: GetParagraphs :many
SELECT id, teacher_id, header, paragraph, created_at
FROM paragraphs;

-- name: GetParagraphByID :one
SELECT id, teacher_id, header, paragraph, created_at
FROM paragraphs
WHERE id = $1;

-- name: GetParagraphsByTeacher :many
SELECT id, teacher_id, header, paragraph, created_at
FROM paragraphs
WHERE teacher_id = $1;

-- name: CreateParagraph :one
INSERT INTO paragraphs (teacher_id, header, paragraph)
VALUES ($1, $2, $3)
RETURNING id, teacher_id, header, paragraph, created_at;

-- name: UpdateParagraph :one
UPDATE paragraphs
SET teacher_id = $1, header = $2, paragraph = $3
WHERE id = $4
RETURNING id, teacher_id, header, paragraph, created_at;

-- name: DeleteParagraph :exec
DELETE FROM paragraphs
WHERE id = $1;
