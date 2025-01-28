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
	log.Println("Request inciada")
	defer log.Println("Request finalizada")

	select {
	case <-time.After(time.Second * 5):
		log.Println("Request processada com sucesso - Cmd")
		w.Write([]byte("Request processada com sucesso - Browser"))
	case <-ctx.Done():
		log.Println("Request cancelada pelo cliente - Cmd")
	}
}
