package repository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/kroma-labs/poker-ledger-be/internal/pkg/logger"
	"github.com/rotisserie/eris"
)

type txKey string

const tx txKey = "repository_tx"

// SqlxExecutor interface that includes all common sqlx methods
type SqlxExecutor interface {
	// Query methods
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row

	// Exec methods
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)

	// Named query methods
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	NamedExec(query string, arg interface{}) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)

	// Get/Select methods
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error

	// Prepared statement methods
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
	PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)

	// Rebind for different SQL dialects
	Rebind(query string) string
	BindNamed(query string, arg interface{}) (string, []interface{}, error)
}

type Transactor interface {
	WithinTransaction(ctx context.Context, serviceFn func(context.Context) error) error
	GetExecutor(ctx context.Context) SqlxExecutor
}

type transactorSqlx struct {
	db *sqlx.DB
}

func NewTransactor(
	db *sqlx.DB,
) Transactor {
	return &transactorSqlx{
		db: db,
	}
}

func (t *transactorSqlx) WithinTransaction(
	ctx context.Context,
	serviceFn func(context.Context) error,
) error {
	ctxValue := ctx.Value(tx)
	if ctxValue != nil {
		return serviceFn(ctx)
	}

	sqlxTx, err := t.db.BeginTxx(ctx, nil)
	if err != nil {
		return eris.Wrap(err, "error starting SQL transaction")
	}
	defer func() {
		txErr := sqlxTx.Rollback()
		if txErr != nil {
			if txErr == sql.ErrTxDone {
				return
			}
			logger.Logger.Error(
				eris.ToString(eris.Wrap(txErr, "error rolling back transaction"), true),
			)
		}
	}()

	txCtx := context.WithValue(ctx, tx, sqlxTx)
	if err = serviceFn(txCtx); err != nil {
		return err
	}

	if err = sqlxTx.Commit(); err != nil {
		return eris.Wrap(err, "error committing transaction")
	}

	return nil
}

// GetExecutor returns either the transaction or the main DB connection
// Both *sqlx.DB and *sqlx.Tx implement the SqlxExecutor interface
func (t *transactorSqlx) GetExecutor(ctx context.Context) SqlxExecutor {
	if txValue := ctx.Value(tx); txValue != nil {
		if sqlxTx, ok := txValue.(*sqlx.Tx); ok {
			return sqlxTx
		}
	}
	return t.db
}
