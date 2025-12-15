-- name: InsertPlayer :one
INSERT INTO players (name, created_at)
VALUES (?, ?)
RETURNING id, name, created_at;
