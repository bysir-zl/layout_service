package tests

import "testing"

func TestTypes(t *testing.T) {
	var a uint = 1
	var b uint = 2

	c := 1<<a | 1<<b

	t.Log(c)

	t.Log(c&(1<<a) == 0)
}
