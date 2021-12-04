package util

import (
	"os"
	"testing"
)

func TestReadInstructions(t *testing.T) {
	sample, err := os.Open("sample")
	if err != nil {
		panic(err)
	}
	defer sample.Close()

	expected := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	actual := ReadFile(sample)

	if len(actual) != len(expected) {
		t.Fatalf("Array length differed, expected %d, actual %d)", len(expected), len(actual))
	}
	for i, entry := range expected {
		if actual[i] != entry {
			t.Errorf("Expected entry %s, got %s", expected, actual[i])
		}
	}
}
