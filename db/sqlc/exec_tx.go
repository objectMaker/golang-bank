package db

import (
	"context"
	"fmt"
)

func (store *SqlStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	//start a transaction
	tx, err := store.connPool.Begin(ctx)
	// if transaction start failed return error
	if err != nil {
		return err
	}
	// transaction is also a query obj , pass it to get a query
	q := New(tx)
	//exec user query
	txErr := fn(q)
	// if user query failed return error
	if txErr != nil {
		rbErr := tx.Rollback(ctx)
		if rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return txErr
	}
	return tx.Commit(ctx)
}
