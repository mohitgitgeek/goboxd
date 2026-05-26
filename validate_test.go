package main

import "testing"

func TestIsValidFilename(t *testing.T) {
tests := []struct {
name     string
filename string
want     bool
}{
{name: "valid", filename: "main.go", want: true},
{name: "invalid absolute path", filename: "/etc/passwd", want: false},
{name: "invalid parent traversal", filename: "../main.go", want: false},
{name: "hidden file", filename: ".env", want: false},
{name: "double dot hidden file", filename: "..env", want: false},
{name: "hidden file in subdirectory", filename: "dir/.hidden", want: false},
}

for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
if got := isValidFilename(tt.filename); got != tt.want {
t.Fatalf("isValidFilename(%q) = %v, want %v", tt.filename, got, tt.want)
}
})
}
}
