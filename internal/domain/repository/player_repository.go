package repository

import (
	"context"

	"github.com/kroma-labs/poker-ledger-be/internal/domain/entity"
	"github.com/rotisserie/eris"
)

type PlayerRepository interface {
	Insert(ctx context.Context, player entity.Player) (entity.Player, error)
}

type playerRepositorySqlx struct {
	t Transactor
}

func NewPlayerRepository(t Transactor) PlayerRepository {
	return &playerRepositorySqlx{t}
}

func (pr *playerRepositorySqlx) Insert(ctx context.Context, player entity.Player) (entity.Player, error) {
	exec := pr.t.GetExecutor(ctx)

	query := `
		INSERT INTO players (name, created_at)
		VALUES (:name, :created_at)
		RETURNING *`

	// Convert named query to positional with args
	boundQuery, args, err := exec.BindNamed(query, player)
	if err != nil {
		return entity.Player{}, eris.Wrap(err, "failed to bind named query")
	}

	// Use context-aware method with positional args
	row := exec.QueryRowxContext(ctx, boundQuery, args...)
	if err := row.StructScan(&player); err != nil {
		return entity.Player{}, eris.Wrap(err, "failed to insert player")
	}

	return player, nil
}
