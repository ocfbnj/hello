package algorithm

import "testing"

func TestNextPermutation(t *testing.T) {
	str := "123"

	NextPermutation(&str)
	if str != "132" {
		t.Errorf("The next permutation of 123 is %v, want 132", str)
	}
}
