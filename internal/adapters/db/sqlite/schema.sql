CREATE TABLE goose_db_version (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		version_id INTEGER NOT NULL,
		is_applied INTEGER NOT NULL,
		tstamp TIMESTAMP DEFAULT (datetime('now'))
	);
CREATE TABLE sqlite_sequence(name,seq);
CREATE TABLE players (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(100) NOT NULL,
    created_at INTEGER NOT NULL
);
CREATE TABLE rooms (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    code CHAR(6) NOT NULL,
    host_player_id INTEGER NOT NULL REFERENCES players(id),
    status VARCHAR(20) NOT NULL DEFAULT 'WAITING',
    config_json TEXT,
    created_at INTEGER NOT NULL
);
CREATE INDEX rooms_code_idx ON rooms (code);
CREATE INDEX rooms_host_player_id_idx ON rooms (host_player_id);
