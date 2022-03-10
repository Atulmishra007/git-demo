package main
import "testing"

func test_add_2_0 (t *testing.T) {
	/ define variables for test / Preparation
	a := 2
	b := 0
	wanted := 2

	// call function to test
	got := Add(a, b)

	// verify result
	if got != wanted {
		t.Errorf("Add(%d,%d) = %d, wanted %d", a, b, got, wanted)
	}

  t.Log("hello")
}