package main

import "testing"

func TestAdd(t *testing.T) {
	res := Add(7, 5)
	if res != 12 {
		t.Errorf("Add(7, 5) = %d; want 12", res)
	}
}

func TestSub(t *testing.T) {
	res := Sub(7, 5)
	if res != 2 {
		t.Errorf("Sub(7, 5) = %d; want 2", res)
	}
}
