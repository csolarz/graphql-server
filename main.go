package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Servidor iniciado en http://localhost:8080/graphql")
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong!")
	})
	http.ListenAndServe(":8080", nil)
}
