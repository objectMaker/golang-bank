package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/objectMaker/golang-bank/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	var a CreateAccountParams
	a.Owner = utils.RandomName()
	a.Balance = 100
	a.Currency = utils.RandomName()
	Account, err := testStore.CreateAccount(context.Background(), a)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(a.Owner)
	require.Equal(t, Account.Owner, a.Owner)
	require.Equal(t, Account.Balance, a.Balance)
	require.Equal(t, Account.Currency, a.Currency)
}
