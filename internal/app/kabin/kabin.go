package kabin

import (
	"flag"
	"net/http"
	"strconv"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Run() {
	flagPort := flag.Int("port", 47623, "the port to host the server on")
	flag.Parse()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)

	r.Get("/", editor)

	fs := http.FileServer(http.Dir(filepath.Join("web", "static")))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":"+strconv.Itoa(*flagPort), r) 
}

func editor(w http.ResponseWriter, r *http.Request) {
	page, err := os.ReadFile(filepath.Join("web", "index.html"))

	if err != nil {
		w.Write([]byte("ERROR"))
		panic(err)
	}

	w.Write(page)
}
