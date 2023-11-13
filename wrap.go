package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var _ DBTX = (*wrappedDB)(nil)

func Wrap(db DBTX) DBTX {
	return &wrappedDB{db}
}

type wrappedDB struct {
	DBTX
}

func (w wrappedDB) ExecContext(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	if b, ok := BuilderFrom(ctx); ok {
		query, args = b.Build(query, args...)
	}

	return w.DBTX.ExecContext(ctx, query, args...)
}

func (w wrappedDB) QueryContext(ctx context.Context, query string, args ...interface{}) (*pgx.Rows, error) {
	if b, ok := BuilderFrom(ctx); ok {
		query, args = b.Build(query, args...)
	}

	return w.DBTX.QueryContext(ctx, query, args...)
}

func (w wrappedDB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *pgx.Row {
	if b, ok := BuilderFrom(ctx); ok {
		query, args = b.Build(query, args...)
	}

	return w.DBTX.QueryRowContext(ctx, query, args...)
}
