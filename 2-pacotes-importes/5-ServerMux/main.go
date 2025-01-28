package main

import "net/http"

type blog struct {
	title string
}

// go possui multiplexer padrão
// podemos criar o nosso multiplexer para ter maior controle
func main() {

	// nosso multiplexer - ServerMux
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "My Blog"})
	http.ListenAndServe(":8080", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
