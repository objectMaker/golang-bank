package db

import (
	"context"
	"fmt"
)

type TransferReq struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// the transferRes
type TransferRes struct {
	Transfer    Transfer `json:"transfer"`
	ToAccount   Account  `json:"to_account"`
	FromAccount Account  `json:"from_account"`
	ToEntry     Entry    `json:"to_entry"`
	FromEntry   Entry    `json:"from_entry"`
}

// define a backgroundWithValueKey , besides the func variable declaration must have the var keyword
var BackgroundWithValueKey = struct{}{}

// create a transfer
func (store *SqlStore) TransferTx(ctx context.Context, req *TransferReq) (TransferRes, error) {
	var transferRes TransferRes
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		//first: create a transfer
		index := ctx.Value(BackgroundWithValueKey)
		fmt.Printf("createTransfer %v \n", index)
		transferRes.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: req.FromAccountID,
			ToAccountID:   req.ToAccountID,
			Amount:        req.Amount,
		})
		if err != nil {
			return err
		}
		//fourth: create to entry
		fmt.Printf("CreateToEntry %v \n", index)
		transferRes.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: req.ToAccountID,
			Amount:    req.Amount,
		})
		if err != nil {
			return err
		}
		//fifth: create from entry
		fmt.Printf("CreateFromEntry %v \n", index)
		transferRes.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: req.FromAccountID,
			Amount:    -req.Amount,
		})
		if err != nil {
			return err
		}
		//second: get to account to update its balance
		fmt.Printf("GetAccountToForUpdate %v \n", index)
		toAccount, err := q.GetAccountForUpdate(ctx, req.ToAccountID)
		if err != nil {
			return err
		}
		fmt.Printf("UpdateToAccount %v \n", index)
		changedToAccount, err := q.UpdateAccount(ctx, UpdateAccountParams{
			ID:      req.ToAccountID,
			Balance: toAccount.Balance + req.Amount,
		})
		if err != nil {
			return err
		}
		transferRes.ToAccount = changedToAccount
		//third: get from account to update its balance
		fmt.Printf("GetAccountFromForUpdate %v \n", index)
		fromAccount, err := q.GetAccountForUpdate(ctx, req.FromAccountID)
		if err != nil {
			return err
		}
		fmt.Printf("UpdateFromAccount %v \n", index)
		changedFromAccount, err := q.UpdateAccount(ctx, UpdateAccountParams{
			ID:      req.FromAccountID,
			Balance: fromAccount.Balance - req.Amount,
		})
		if err != nil {
			return err
		}
		transferRes.FromAccount = changedFromAccount
		//if everything fine return nil
		return nil
	})
	if err != nil {
		return transferRes, err
	} else {
		return transferRes, nil
	}
}
