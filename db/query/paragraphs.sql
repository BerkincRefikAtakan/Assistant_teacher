-- name: GetParagraphs :one
SELECT id, teacher_id, header, paragraph, created_at
FROM paragraphs
WHERE teacher_id = $1 AND header=$2;

-- name: GetParagraphsByTeacher :many
SELECT id, teacher_id, header, paragraph, created_at
FROM paragraphs
WHERE teacher_id = $1
LIMIT $2 OFFSET $3;

-- name: CreateParagraph :one
INSERT INTO paragraphs (teacher_id, header, paragraph)
VALUES ($1, $2, $3)
RETURNING id, teacher_id, header, paragraph, created_at;

-- name: UpdateParagraphOrAndHeader :one
UPDATE paragraphs
SET  header = $1, paragraph = $2
WHERE id = $3
RETURNING id, teacher_id, header, paragraph, created_at;

-- name: DeleteParagraph :exec
DELETE FROM paragraphs
WHERE teacher_id = $1 AND header=$2;
