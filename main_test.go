package main

import "testing"

func Test_add(t *testing.T) {

	expected := 2 + 2
	got := add(2, 2)

	if expected != got {
		t.Fatal("Failed")
	}

}
