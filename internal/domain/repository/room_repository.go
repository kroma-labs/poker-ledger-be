package repository

import (
	"context"

	"github.com/kroma-labs/poker-ledger-be/internal/domain/entity"
	"github.com/rotisserie/eris"
)

type RoomRepository interface {
	Insert(ctx context.Context, room entity.Room) (entity.Room, error)
	GetAll(ctx context.Context) ([]entity.Room, error)
}

type roomRepositorySqlx struct {
	t Transactor
}

func NewRoomRepository(t Transactor) RoomRepository {
	return &roomRepositorySqlx{t}
}

func (rr *roomRepositorySqlx) Insert(ctx context.Context, room entity.Room) (entity.Room, error) {
	query := `
		INSERT INTO rooms (code, host_player_id, status, config_json, created_at)
		VALUES (:code, :host_player_id, :status, :config_json, :created_at)
		RETURNING *`

	exec := rr.t.GetExecutor(ctx)

	boundQuery, args, err := exec.BindNamed(query, room)
	if err != nil {
		return entity.Room{}, eris.Wrap(err, "failed to bind named query")
	}

	row := exec.QueryRowxContext(ctx, boundQuery, args...)
	if err := row.StructScan(&room); err != nil {
		return entity.Room{}, eris.Wrap(err, "failed to insert room")
	}

	return room, nil
}

func (rr *roomRepositorySqlx) GetAll(ctx context.Context) ([]entity.Room, error) {
	query := `SELECT * FROM rooms`

	exec := rr.t.GetExecutor(ctx)
	var rooms []entity.Room
	if err := exec.SelectContext(ctx, &rooms, query); err != nil {
		return nil, eris.Wrap(err, "error querying rooms")
	}

	return rooms, nil
}
