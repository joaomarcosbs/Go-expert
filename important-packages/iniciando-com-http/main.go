package main

import "net/http"

func BuscaCEP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Busca CEP"))
}

func main() {
	http.HandleFunc("/", BuscaCEP)
	http.ListenAndServe(":8080", nil)
}
