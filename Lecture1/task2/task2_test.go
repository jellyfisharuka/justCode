package main

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	t.Run("Empty Input", func(t *testing.T) {
		input := []string{}
		result := longestCommonPrefix(input)
		expected := ""
		if result != expected {
			t.Errorf("Expected '%s', but got '%s'", expected, result)
		}
	})

	t.Run("No Common Prefix", func(t *testing.T) {
		input := []string{"apple", "banana", "cherry"}
		result := longestCommonPrefix(input)
		expected := ""
		if result != expected {
			t.Errorf("Expected '%s', but got '%s'", expected, result)
		}
	})

	t.Run("Common Prefix Exists", func(t *testing.T) {
		input := []string{"flower", "flow", "flour"}
		result := longestCommonPrefix(input)
		expected := "flo"
		if result != expected {
			t.Errorf("Expected '%s', but got '%s'", expected, result)
		}
	})

	t.Run("Single Word Input", func(t *testing.T) {
		input := []string{"dog"}
		result := longestCommonPrefix(input)
		expected := "dog"
		if result != expected {
			t.Errorf("Expected '%s', but got '%s'", expected, result)
		}
	})
}
