package main

import (
"net/http"
"path/filepath"
"strings"
)

func isValidFilename(name string) bool {
if name == "" || strings.ContainsRune(name, '\x00') {
return false
}
if filepath.IsAbs(name) {
return false
}
clean := filepath.Clean(name)
if clean == "." || clean == ".." || strings.HasPrefix(clean, ".."+string(filepath.Separator)) {
return false
}
if strings.HasPrefix(filepath.Base(clean), ".") {
return false
}
return clean == name
}

func truncateOutput(output string, maxBytes int) string {
if maxBytes <= 0 {
return ""
}
if len(output) <= maxBytes {
return output
}
return output[:maxBytes]
}

func main() {
mux := http.NewServeMux()
mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
w.WriteHeader(http.StatusOK)
_, _ = w.Write([]byte("ok"))
})

server := &http.Server{
Addr:    ":8080",
Handler: mux,
}
if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
panic(err)
}
}
