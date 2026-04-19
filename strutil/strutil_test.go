package strutil

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		input, want string
	}{
		{"hello", "olleh"},
		{"ab", "ba"},
		{"a", "a"},
		{"", ""},
	}
	for _, tc := range tests {
		got := Reverse(tc.input)
		if got != tc.want {
			t.Errorf("Reverse(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"hello world", 2},
		{"one", 1},
		{"", 0},
		{"  spaces  everywhere  ", 2},
	}
	for _, tc := range tests {
		got := CountWords(tc.input)
		if got != tc.want {
			t.Errorf("CountWords(%q) = %d, want %d", tc.input, got, tc.want)
		}
	}
}
