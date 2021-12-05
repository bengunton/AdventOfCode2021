package main

import "testing"

func TestParseLine(t *testing.T) {
	cases := []struct {
		input     string
		expectErr bool
		expected  Line
	}{
		{
			input:     `1,2 -> 1,4`,
			expectErr: false,
			expected:  Line{Point{1, 2}, Point{1, 4}},
		},
		{
			input:     `1,a2 -> 1,4`,
			expectErr: true,
			expected:  Line{},
		},
		{
			input:     `111,222 -> 111,444`,
			expectErr: false,
			expected:  Line{Point{111, 222}, Point{111, 444}},
		},
	}

	for _, test := range cases {
		actual, err := parseLine(test.input)
		if test.expectErr && err == nil {
			t.Errorf("Expected error but got none")
		}
		if !test.expectErr && err != nil {
			t.Errorf("Expected no error but got %v", err)
		}
		if actual != test.expected {
			t.Errorf("Failed to parse line. Expected %v, got %v", test.expected, actual)
		}
	}
}
