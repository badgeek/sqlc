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

func (w wrappedDB) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	if b, ok := BuilderFrom(ctx); ok {
		query, args = b.Build(query, args...)
	}

	return w.DBTX.Exec(ctx, query, args...)
}

func (w wrappedDB) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	if b, ok := BuilderFrom(ctx); ok {
		query, args = b.Build(query, args...)
	}

	return w.DBTX.Query(ctx, query, args...)
}

func (w wrappedDB) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	if b, ok := BuilderFrom(ctx); ok {
		query, args = b.Build(query, args...)
	}

	return w.DBTX.QueryRow(ctx, query, args...)
}
