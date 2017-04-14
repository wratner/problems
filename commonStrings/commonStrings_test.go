package main

import (
	"strings"
	"testing"
)

func TestGetLongestCommonSubstring(t *testing.T) {
	expected := " is awesome!"
	result := getLongestCommonSubstring("Everything is awesome!", "Hello World is awesome!")
	if result != expected {
		t.Error("Expected: is awesome!, Received: ", result)
	}
}

func TestReadInput(t *testing.T) {
	input := "test\n"
	result, _ := readInput(strings.NewReader(input))
	if result != strings.TrimSpace(input) {
		t.Error("Expected: " + input + " Received: " + result)
	}
}
