package main

import (
	"fmt"
	"net/http"

	"github.com/rsacramento/goServer/internal/routes"
)

func main() {
	router := routes.NewRouter()

	port := 8086
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server running on http://localhost:%d\n", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
