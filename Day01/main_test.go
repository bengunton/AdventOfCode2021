package main

import (
	"os"
	"testing"
)

func TestCountDepthIncreases(t *testing.T) {
	test, err := os.Open("./sample")
	if err != nil {
		panic(err)
	}
	defer test.Close()

	numIncreases := CountDepthIncreases(test)
	if numIncreases != 7 {
		t.Errorf("Expected 7 increases, got %d", numIncreases)
	}
}
