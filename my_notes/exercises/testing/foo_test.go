package mymath

import "testing"

func TestAdd(t *testing.T) {
	a, b := 1, 2
	v := Add(a, b)
	exp := 4
	if v != exp {
		t.Errorf("Expected %v, got %v", exp, v)
	}
}
