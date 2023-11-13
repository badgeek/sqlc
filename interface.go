package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	QueryContext(context.Context, string, ...interface{}) (*pgx.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *pgx.Row
}
