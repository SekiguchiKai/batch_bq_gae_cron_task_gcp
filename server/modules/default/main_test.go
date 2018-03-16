package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("2 + 3 = 5になること", func(t *testing.T) {
		const Wants =  5
		actual := sum(2, 3)
		if actual != Wants {
			t.Fatalf("actual = %d, wants = %d", actual, Wants)
		}
	})
}