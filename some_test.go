package gonat

import (
	"testing"
)

func TestCalc(t *testing.T) {
	t.Run("add numbers", func(t *testing.T) {
		got := add(1, 5)
		expected := 6
		if got != expected {
			t.Errorf("Got: %v, Expected %v", got, expected)
		}
	})

	t.Run("add other numbers ", func(t *testing.T) {
		got := add(45, 17)
		expected := 62
		if got != expected {
			t.Errorf("Got: %v, Expected %v", got, expected)
		}
	})

	t.Run("add negative numbers ", func(t *testing.T) {
		got := add(-45, -17)
		expected := -62
		if got != expected {
			t.Errorf("Got: %v, Expected %v", got, expected)
		}
	})

	t.Run("add zero numbers ", func(t *testing.T) {
		got := add(0, 0)
		expected := 0
		if got != expected {
			t.Errorf("Got: %v, Expected %v", got, expected)
		}
	})
}
