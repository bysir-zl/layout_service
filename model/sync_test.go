package model

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"container/heap"
)

func TestSync(t *testing.T) {
	err := SyncAll()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")

	heap.Fix()
}
