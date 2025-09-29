package main

import (
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")

	select {
	case <-time.After(5 * time.Second):
		// imprime na linha de comando
		log.Println("Request processada com sucesso")
		// imprime no navegador
		w.Write([]byte("Request processada com sucesso"))
	case <-ctx.Done():
		// imprime na linha de comando
		log.Println("Request cancelada pelo cliente")
	}
}

func main() {
	http.ListenAndServe(":8080", nil)
	http.HandleFunc("/", handler)
}
