package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/kroma-labs/poker-ledger-be/internal/adapters/db/sqlite/sqlc"
	"github.com/kroma-labs/poker-ledger-be/internal/domain/entity"
	"github.com/rotisserie/eris"
)

type RoomRepository interface {
	Insert(ctx context.Context, room entity.Room) (entity.Room, error)
}

type roomRepositorySqlc struct {
	querier sqlc.Querier
}

func NewRoomRepository(db *sql.DB) *roomRepositorySqlc {
	return &roomRepositorySqlc{
		sqlc.New(db),
	}
}

func (rr *roomRepositorySqlc) Insert(ctx context.Context, room entity.Room) (entity.Room, error) {
	querier, err := rr.getQuerier(ctx)
	if err != nil {
		return entity.Room{}, err
	}

	configJSON := sql.NullString{}
	if room.ConfigJSON != "" {
		configJSON.String = room.ConfigJSON
		configJSON.Valid = true
	}

	insertedRoom, err := querier.InsertRoom(ctx, sqlc.InsertRoomParams{
		Code:         room.Code,
		HostPlayerID: int64(room.HostPlayerID),
		Status:       string(room.Status),
		ConfigJson:   configJSON,
		CreatedAt:    room.CreatedAt.Unix(),
	})
	if err != nil {
		return entity.Room{}, eris.Wrap(err, "error inserting room")
	}

	code, ok := insertedRoom.Code.(string)
	if !ok {
		return entity.Room{}, eris.Errorf("error asserting code as string, code is: %T", insertedRoom.Code)
	}

	return entity.Room{
		ID:           int(insertedRoom.ID),
		Code:         code,
		HostPlayerID: int(insertedRoom.HostPlayerID),
		Status:       entity.RoomStatus(insertedRoom.Status),
		ConfigJSON:   insertedRoom.ConfigJson.String,
		CreatedAt:    time.Unix(insertedRoom.CreatedAt, 0),
	}, nil
}

func (rr *roomRepositorySqlc) getQuerier(ctx context.Context) (sqlc.Querier, error) {
	ctxValue := ctx.Value(tx)
	if ctxValue == nil {
		return rr.querier, nil
	}

	queries, ok := ctxValue.(sqlc.Querier)
	if !ok {
		return nil, eris.Errorf("error asserting tx as sqlc.Querier. tx is: %T", ctxValue)
	}

	return queries, nil
}
