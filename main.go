package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	server := http.Server{
		Addr: ":3000",
		Handler: router,
	}

	fmt.Println("Listening on PORT :3000")
	server.ListenAndServe()
}