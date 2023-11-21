package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/objectMaker/golang-bank/utils"
	"github.com/stretchr/testify/require"

	"github.com/jackc/pgx/v5"
)

var b *Queries

func TestCreateAccount(t *testing.T) {
	conn, err := pgx.Connect(context.Background(), "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable")
	if err != nil {
		t.Error(err)
	}
	b = New(conn)
	defer conn.Close(context.Background())
	var a CreateAccountParams
	a.Owner = utils.RandomName()
	a.Balance = 100
	a.Currency = utils.RandomName()
	Account, err := b.CreateAccount(context.Background(), a)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(a.Owner)
	require.Equal(t, Account.Owner, a.Owner)
	require.Equal(t, Account.Balance, a.Balance)
	require.Equal(t, Account.Currency, a.Currency)
}
