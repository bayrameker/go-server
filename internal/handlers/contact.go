// internal/handlers/contact.go

package handlers

import (
    "fmt"
    "net/http"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "İletişim")
}
