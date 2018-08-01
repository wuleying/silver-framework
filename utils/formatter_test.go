package utils

import "testing"

func TestReadableFloat(t *testing.T) {
	var tests = []struct {
		// Test table
		in  float64
		out string
	}{
		{1.001, "1.001"},
		{100.01, "100.01"},
		{1000.99, "1000.99"},
	}

	for i, tt := range tests {
		s := ReadableFloat(tt.in)
		if s != tt.out {
			t.Errorf("%d. %f => %q, wanted: %q", i, tt.in, s, tt.out)
		}
	}
}
