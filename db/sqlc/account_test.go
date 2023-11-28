package db

import (
	"context"
	"log"
	"testing"

	"github.com/objectMaker/golang-bank/utils"
	"github.com/stretchr/testify/require"
)

func createTestAccount(t *testing.T) Account {
	inputAccount := CreateAccountParams{
		Owner:    utils.RandomName(),
		Balance:  utils.RandomInt(1, 100),
		Currency: utils.RandomCurrency(),
	}
	newAccount, err := testStore.CreateAccount(context.Background(), inputAccount)

	if err != nil {
		log.Fatal(err)
	}
	require.Equal(t, inputAccount.Owner, newAccount.Owner)
	require.Equal(t, inputAccount.Balance, newAccount.Balance)
	require.Equal(t, inputAccount.Currency, newAccount.Currency)
	return newAccount
}

func TestCreateAccount(t *testing.T) {
	createTestAccount(t)
}

func TestGetAccount(t *testing.T) {
	//create new account
	newAccount := createTestAccount(t)
	//query the new account by account id
	queryAccount, err := testStore.GetAccount(context.Background(), newAccount.ID)
	if err != nil {
		log.Fatal("get account error", err)
	}
	require.Equal(t, newAccount.ID, queryAccount.ID)
	require.Equal(t, newAccount.Owner, queryAccount.Owner)
	require.Equal(t, newAccount.Balance, queryAccount.Balance)
	require.Equal(t, newAccount.CreatedAt, queryAccount.CreatedAt)
	require.Equal(t, newAccount.Currency, queryAccount.Currency)
}

func TestListAccounts(t *testing.T) {
	arg := ListAccountsParams{
		Owner:  "",
		Limit:  5,
		Offset: 0,
	}
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = createTestAccount(t)
	}
	arg.Owner = lastAccount.Owner
	accounts, err := testStore.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	for _, account := range accounts {
		require.Equal(t, account.Owner, lastAccount.Owner)
		require.NotEmpty(t, account.ID)
	}
}

func TestUpdateAccount(t *testing.T) {
	newAccount := createTestAccount(t)
	newBalance := utils.RandomInt(1, 100)
	arg := UpdateAccountParams{
		ID:      newAccount.ID,
		Balance: newBalance,
	}
	updateAccount, err := testStore.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.Balance, updateAccount.Balance)
}
