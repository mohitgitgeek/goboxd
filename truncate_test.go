package main

import "testing"

func TestTruncateOutput(t *testing.T) {
const input = "0123456789abcdef"
if got := truncateOutput(input, 8); got != "01234567" {
t.Fatalf("truncateOutput() = %q, want %q", got, "01234567")
}
}
