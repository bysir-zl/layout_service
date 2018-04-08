package model

import "testing"

func TestGetLayoutPage(t *testing.T) {
	i,err:=GetLayoutPage(1)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v",i)
}