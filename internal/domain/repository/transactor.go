package repository

import (
	"context"
	"database/sql"

	"github.com/kroma-labs/poker-ledger-be/internal/adapters/db/sqlite/sqlc"
	"github.com/kroma-labs/poker-ledger-be/internal/pkg/logger"
	"github.com/rotisserie/eris"
)

type txKey string

const tx txKey = "repository_tx"

type Transactor interface {
	WithinTransaction(ctx context.Context, serviceFn func(context.Context) error) error
}

type transactorSqlc struct {
	db      *sql.DB
	queries *sqlc.Queries
}

func NewTransactor(
	db *sql.DB,
) Transactor {
	return &transactorSqlc{
		db,
		sqlc.New(db),
	}
}

func (t *transactorSqlc) WithinTransaction(
	ctx context.Context,
	serviceFn func(context.Context) error,
) error {
	ctxValue := ctx.Value(tx)
	if ctxValue != nil {
		return serviceFn(ctx)
	}

	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return eris.Wrap(err, "error starting SQL transaction")
	}
	defer func() {
		txErr := tx.Rollback()
		if txErr != nil {
			if txErr.Error() == "sql: transaction has already been committed or rolled back" {
				return
			}
			logger.Logger.Error(
				eris.ToString(eris.Wrap(txErr, "error rolling back transaction"), true),
			)
		}
	}()

	txQueries := t.queries.WithTx(tx)
	txCtx := context.WithValue(ctx, tx, txQueries)
	if err = serviceFn(txCtx); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return eris.Wrap(err, "error committing transaction")
	}

	return nil
}
