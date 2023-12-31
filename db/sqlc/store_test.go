package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	//create two accounts
	fromAccount := CreateTestAccount(t)
	toAccount := CreateTestAccount(t)
	amount := int64(10)
	//create two channels to receive goroutines
	transferResChannel := make(chan TransferRes, 5)
	errChannel := make(chan error, 5)
	//use the transferTx
	//start a for loop to concurrency transaction
	loopCount := 2
	for i := 0; i < loopCount; i++ {
		// start go routine
		currentIndex := i + 1
		go func() {
			ctxWithValue := context.WithValue(context.Background(), BackgroundWithValueKey, currentIndex)
			transferRes, err := testStore.TransferTx(ctxWithValue, &TransferReq{
				FromAccountID: fromAccount.ID,
				ToAccountID:   toAccount.ID,
				Amount:        amount,
			})
			// if has result , that indicate this go routine should give the value to main thread
			transferResChannel <- transferRes
			errChannel <- err
		}()
	}
	for i := 0; i < loopCount; i++ {
		currentLoopCount := int64(i + 1)
		transferRes := <-transferResChannel
		err := <-errChannel
		require.NoError(t, err)
		//judge the transferRes is valid
		// transfer
		require.Equal(t, fromAccount.ID, transferRes.Transfer.FromAccountID)
		require.Equal(t, toAccount.ID, transferRes.Transfer.ToAccountID)
		require.Equal(t, amount, transferRes.Transfer.Amount)
		// to entry
		require.Equal(t, toAccount.ID, transferRes.ToEntry.AccountID)
		require.Equal(t, amount, transferRes.ToEntry.Amount)
		//from entry
		require.Equal(t, fromAccount.ID, transferRes.FromEntry.AccountID)
		require.Equal(t, amount, -transferRes.FromEntry.Amount)
		//test account balance
		require.NotEmpty(t, transferRes.ToAccount)
		require.NotEmpty(t, transferRes.FromAccount)
		fmt.Printf("ToAccountBalance %v %v \n", i+1, transferRes.ToAccount.Balance)
		fmt.Printf("FromAccountBalance %v %v \n", i+1, transferRes.FromAccount.Balance)
		require.Equal(t, toAccount.Balance+currentLoopCount*amount, transferRes.ToAccount.Balance)
		require.Equal(t, fromAccount.Balance-currentLoopCount*amount, transferRes.FromAccount.Balance)
	}
}
