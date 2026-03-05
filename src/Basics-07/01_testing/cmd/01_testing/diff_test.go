package main

import "testing"

func TestDiff(t *testing.T) {
	if diff(10, 3) != 7 {
		t.Fatal("fail")
	}
}
