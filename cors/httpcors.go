package main

import (
	"log"
	"net/http"
)

func httpCors() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "https://www.google.com")
		w.Header().Set("Access-Control-Allow-Method", "OPTIONS, GET, POST, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")

		if r.Method == "OPTIONS" {
			w.Write([]byte("Allowed"))
			return
		}

		w.Write([]byte("Hello"))
	})
	log.Println("Server is running at http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
