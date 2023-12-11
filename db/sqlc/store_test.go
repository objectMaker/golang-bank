package db

import (
	"context"
	"testing"
)

func TestTransferTx(t *testing.T) {
	//create two accounts
	fromAccount := CreateTestAccount(t)
	toAccount := CreateTestAccount(t)
	amount := int64(10)
	//use the transferTx
	transferRes, err := testStore.TransferTx(context.Background(), &TransferReq{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Amount:        amount,
	})
	if err != nil {
		t.Fatal(err)
	}
	// transferRes
}
