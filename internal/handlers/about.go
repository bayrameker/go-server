// internal/handlers/about.go

package handlers

import (
    "fmt"
    "net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hakkımızda")
}
