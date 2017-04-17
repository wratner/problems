package main

import (
	"reflect"
	"testing"
)

func TestGetSequences(t *testing.T) {
	expected := []string{"XMJYAUZ;MZJAWXU", "AATCC;ACACG", "invalidSequence"}
	sequences, err := getSequences("test.txt")
	if err != nil {
		t.Error("Error: ", err.Error)
	}
	if !reflect.DeepEqual(expected, sequences) {
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
	if !reflect.DeepEqual(expected, subsequences) {
		t.Errorf("Expected: %v Received: %v", expected, subsequences)
	}
}
