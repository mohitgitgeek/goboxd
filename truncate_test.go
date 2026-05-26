package main

import "testing"

func TestTruncateOutput(t *testing.T) {
const input = "0123456789abcdef"
tests := []struct {
name     string
input    string
maxBytes int
want     string
}{
{name: "truncated", input: input, maxBytes: 8, want: "01234567"},
{name: "shorter than max", input: "abc", maxBytes: 8, want: "abc"},
{name: "zero max", input: input, maxBytes: 0, want: ""},
{name: "negative max", input: input, maxBytes: -1, want: ""},
{name: "empty input", input: "", maxBytes: 8, want: ""},
}

for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
	if got := truncateOutput(tt.input, tt.maxBytes); got != tt.want {
		t.Fatalf("truncateOutput() = %q, want %q", got, tt.want)
	}
})
}
}
