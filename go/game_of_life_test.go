package main

import (
	"testing"
)

func TestTestFuncReturnsFoo(t *testing.T) {
	result := testFunc()

	if result != "foo" {
		t.Error("testFunc did not return 'foo'")
	}
}
