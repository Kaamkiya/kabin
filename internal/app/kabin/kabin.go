package kabin

import (
	"flag"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Run() {
	flagPort := flag.Int("port", 47623, "the port to host the server on")
	flag.Parse()

	r := chi.NewRouter()
	r.Get("/", editor)

	http.ListenAndServe(":"+strconv.Itoa(*flagPort), r) 
}

func editor(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
