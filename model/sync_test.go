package model

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestSync(t *testing.T) {
	err := SyncAll()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}
