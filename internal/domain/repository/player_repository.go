package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/kroma-labs/poker-ledger-be/internal/adapters/db/sqlite/sqlc"
	"github.com/kroma-labs/poker-ledger-be/internal/domain/entity"
	"github.com/rotisserie/eris"
)

type PlayerRepository interface {
	Insert(ctx context.Context, player entity.Player) (entity.Player, error)
}

type playerRepositorySqlc struct {
	querier sqlc.Querier
}

func NewPlayerRepository(db *sql.DB) *playerRepositorySqlc {
	return &playerRepositorySqlc{
		sqlc.New(db),
	}
}

func (pr *playerRepositorySqlc) Insert(ctx context.Context, player entity.Player) (entity.Player, error) {
	querier, err := pr.getQuerier(ctx)
	if err != nil {
		return entity.Player{}, err
	}

	insertedPlayer, err := querier.InsertPlayer(ctx, sqlc.InsertPlayerParams{
		Name:      player.Name,
		CreatedAt: player.CreatedAt.Unix(),
	})
	if err != nil {
		return entity.Player{}, eris.Wrap(err, "error inserting player")
	}

	return entity.Player{
		ID:        int(insertedPlayer.ID),
		Name:      insertedPlayer.Name,
		CreatedAt: time.Unix(insertedPlayer.CreatedAt, 0),
	}, nil
}

func (pr *playerRepositorySqlc) getQuerier(ctx context.Context) (sqlc.Querier, error) {
	ctxValue := ctx.Value(tx)
	if ctxValue == nil {
		return pr.querier, nil
	}

	queries, ok := ctxValue.(sqlc.Querier)
	if !ok {
		return nil, eris.Errorf("error asserting tx as sqlc.Querier. tx is: %T", ctxValue)
	}

	return queries, nil
}
