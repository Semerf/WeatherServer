package server

import (
	"fmt"
	"net/http"
)

func HandlerAll(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	switch r.Method {
	case http.MethodGet:

		w.Header().Set("Content-type", "text/html")
		w.WriteHeader(http.StatusOK)

		println("Hello from Go server")

	default:
		println("default")
	}
}

func Server() {
	fmt.Println("Server start...")

	http.HandleFunc("/", HandlerAll)

	http.ListenAndServe(":8081", nil)
}
