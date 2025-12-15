-- name: InsertRoom :one
INSERT INTO rooms (code, host_player_id, status, config_json, created_at)
VALUES (?, ?, ?, ?, ?)
RETURNING id, code, host_player_id, status, config_json, created_at;
