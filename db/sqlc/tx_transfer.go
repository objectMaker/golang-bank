package db

import "context"

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

// create a transfer
func (store *SqlStore) TransferTx(ctx context.Context, req *TransferReq) (TransferRes, error) {
	var transferRes TransferRes
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		//first: create a transfer
		transferRes.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: req.FromAccountID,
			ToAccountID:   req.ToAccountID,
			Amount:        req.Amount,
		})
		if err != nil {
			return err
		}
		//fourth: create to entry
		transferRes.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: req.ToAccountID,
			Amount:    req.Amount,
		})
		if err != nil {
			return err
		}
		//fifth: create from entry
		transferRes.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: req.FromAccountID,
			Amount:    -req.Amount,
		})
		if err != nil {
			return err
		}
		//TODO: change user balance
		// //second: get to account to update its balance
		// toAccount, err := q.GetAccount(ctx, req.ToAccountID)
		// if err != nil {
		// 	return err
		// }
		// transferRes.ToAccount = toAccount
		// //third: get from account to update its balance
		// fromAccount, err := q.GetAccount(ctx, req.FromAccountID)
		// if err != nil {
		// 	return err
		// }
		// transferRes.FromAccount = fromAccount

		//if everything fine return nil
		return nil
	})
	if err != nil {
		return transferRes, err
	} else {
		return transferRes, nil
	}
}
