package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

//定义store类型

type Store interface {
	Querier
}

var conn *pgx.Conn

func newTestStore() Store {
	var err error
	conn, err = pgx.Connect(context.Background(), "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable")
	if err != nil {
		log.Fatal("connect error: ", err)
	}
	//添加db

	return New(conn)
}
