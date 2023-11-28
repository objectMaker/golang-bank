package db

import (
	"context"
	"os"
	"testing"
)

var testStore Store

func TestMain(m *testing.M) {
	//create testStore for all testFunc
	testStore = newTestStore()
	// after testing close the connection
	defer conn.Close(context.Background())
	//run testCode
	code := m.Run()
	//after testCode
	os.Exit(code)
}
