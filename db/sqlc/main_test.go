package db

import (
	"context"
	"os"
	"log"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store

func TestMain(m *testing.M) {
	//create testStore for all testFunc
	conn, err := pgxpool.New(context.Background(), "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable")
	if err != nil{
		log.Fatal(err)
	}
	testStore = newStore(conn)
	// after testing close the connection
	defer conn.Close()
	//run testCode
	code := m.Run()
	//after testCode
	os.Exit(code)
}
