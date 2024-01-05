// C:\Users\bayra\OneDrive\Masaüstü\GO\backend\cmd\server\main.go

package main

import (
	"fmt"
	"log"
	"net/http"

	"backend/internal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	http.HandleFunc("/contact", handlers.ContactHandler)

	fmt.Println("Server starting on port :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
