package utils

import "testing"

func TestStringSub(t *testing.T) {
	var tests = []struct {
		// Test table
		in  string
		out string
	}{
		{"hello world", "hel"},
		{"你好世界", "你好世"},
	}

	for i, tt := range tests {
		s := StringSub(tt.in, 0, 3)
		if s != tt.out {
			t.Errorf("%d. %s => %q, wanted: %q", i, tt.in, s, tt.out)
		}
	}
}
