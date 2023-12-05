package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

//定义store类型

type Store interface {
	Querier
	execTx(ctx context.Context, fn func(*Queries) error) error
}
type SqlStore struct {
	connPool *pgxpool.Pool
	*Queries
}

func newStore(conn *pgxpool.Pool) Store {
	return &SqlStore{connPool: conn,
		Queries: New(conn),
	}
}
