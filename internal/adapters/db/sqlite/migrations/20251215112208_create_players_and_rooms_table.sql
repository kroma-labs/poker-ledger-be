-- +goose Up
-- +goose StatementBegin
CREATE TABLE players (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE rooms (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    code CHAR(6) NOT NULL,
    host_player_id INTEGER NOT NULL REFERENCES players(id),
    status VARCHAR(20) NOT NULL DEFAULT 'WAITING',
    config_json TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX rooms_code_idx ON rooms (code);
CREATE INDEX rooms_host_player_id_idx ON rooms (host_player_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX rooms_host_player_id_idx;
DROP INDEX rooms_code_idx;
DROP TABLE rooms;
DROP TABLE players;
-- +goose StatementEnd
