package db

import (
	"context"
	"testing"

	"github.com/objectMaker/golang-bank/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	var inputAccount CreateAccountParams
	inputAccount.Owner = utils.RandomName()
	inputAccount.Balance = utils.RandomInt(1, 100)
	inputAccount.Currency = utils.RandomCurrency()

	newAccount, err := testStore.CreateAccount(context.Background(), inputAccount)

	if err != nil {
		t.Error(err)
	}

	require.Equal(t, inputAccount.Owner, newAccount.Owner)
	require.Equal(t, inputAccount.Balance, newAccount.Balance)
	require.Equal(t, inputAccount.Currency, newAccount.Currency)
}
