package main
import "testing"

func Test_add_2_0 (t *testing.T) {

	t.Log("hello")

	// define variables for test / Preparation
	
	a := 2
	b := 0
	wanted := 2

	// call function to test
	got := Add(a, b)

	// verify result
	if got != wanted {
		t.Errorf("Add(%d,%d) = %d, wanted %d", a, b, got, wanted)
	}

  
}
func Test_sub_2_3 (t *testing.T) {

	t.Log("hello")

	// define variables for test / Preparation
	
	a := 2
	b := 3
	wanted :=-1 

	// call function to test
	got := Sub(a, b)

	// verify result
	if got != wanted {
		t.Errorf("sub(%d,%d) = %d, wanted %d", a, b, got, wanted)
	}

 
}