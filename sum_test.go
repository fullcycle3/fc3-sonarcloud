package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	result := sum(2, 3)

	if result != 5 {
		t.Error("O resultado tem que ser 5")
	}
}
