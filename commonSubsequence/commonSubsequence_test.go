package main

import "testing"

func TestGetSequences(t *testing.T) {
	expected := []string{"XMJYAUZ;MZJAWXU", "AATCC;ACACG", "invalidSequence"}
	sequences, err := getSequences("test.txt")
	if err != nil {
		t.Error("Error: ", err.Error)
	}
	if expected[0] != sequences[0] ||
		expected[1] != sequences[1] ||
		expected[2] != sequences[2] {
		t.Errorf("Expected: %v Received: %v", expected, sequences)
	}
}

func TestInvalidFile(t *testing.T) {
	_, err := getSequences("garbage.abc")
	if err == nil {
		t.Error("Expected Error")
	}
}

func TestGetSubSequences(t *testing.T) {
	sequences := []string{"XMJYAUZ;MZJAWXU", "AATCC;ACACG", "invalidSequence"}
	expected := []string{"MJAU", "AAC"}

	subsequences, err := getSubsequences(sequences)
	if err != nil {
		t.Error("Error: ", err.Error)
	}
	if expected[0] != subsequences[0] ||
		expected[1] != subsequences[1] {
		t.Errorf("Expected: %v Received: %v", expected, subsequences)
	}
}
