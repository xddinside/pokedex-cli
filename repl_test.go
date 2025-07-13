package main

import (
	"testing"
)


func TestCleanInput (t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "   XDD egoisto   ",
			expected: []string{"xdd", "egoisto"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		expected := c.expected
		if len(actual) != len(expected) {
			t.Errorf("expected: %v \ngot input: %v", expected, actual)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := expected[i]
			if word != expectedWord {
				t.Errorf("expected: %v \ngot input: %v", expected, actual)
			}
		}
	}
}
