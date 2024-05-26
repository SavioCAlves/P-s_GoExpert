package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciada.")
	defer log.Printf("Request finalizada.")

	select {
	case <-time.After(5 * time.Second):
		// Imprime do prompt do servidor
		log.Println("Request processada com sucesso.")
		// Impreme do browser
		w.Write([]byte("Request processada com sucesso."))
	case <-ctx.Done():
		log.Println("Request cancelada pelo cliente.go")
	}
}
