package model

import (
	"testing"
	"encoding/json"
)

func TestGetUserByIc(t *testing.T) {
	u,err:=GetUserByIc("1")
	if err != nil {
		t.Fatal(err)
	}

	bs,_:=json.MarshalIndent(u,"","  ")
	t.Log(string(bs))
}
