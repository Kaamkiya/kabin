package kabin

import (
	"flag"
	"net/http"
	"strconv"
	"encoding/json"
	"log/slog"
	"os"
	"path/filepath"
	"html/template"

	"codeberg.org/Kaamkiya/kabin/internal/pkg/db"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var tmpl *template.Template

func Run() {
	// Flags.
	flagPort := flag.Int("port", 47623, "the port to host the server on")
	flag.Parse()

	// Database setup.
	if err := db.Open(); err != nil {
		slog.Error("Failed to open database")
		os.Exit(1)
	}
	defer db.Close()

	if err := db.Init(); err != nil {
		slog.Error("Failed to initialize database")
		os.Exit(1)
	}

	// Templating.
	tmplText, err := os.ReadFile(filepath.Join("web", "index.tmpl"))
	if err != nil {
		slog.Error("Failed to read template")
		os.Exit(1)
	}
	tmpl = template.Must(template.New("editor").Parse(string(tmplText)))

	// Routing and HTTP setup.
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)

	r.Get("/", editor)
	r.Get("/{pasteID}", getPaste)
	r.Post("/save", save)

	fs := http.FileServer(http.Dir(filepath.Join("web", "static")))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":"+strconv.Itoa(*flagPort), r) 
}

func editor(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.Execute(w, db.Paste{}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ERROR 500"))
	}
}

func getPaste(w http.ResponseWriter, r *http.Request) {
	paste, err := db.GetPaste(chi.URLParam(r, "pasteID"))
	if err != nil {
		w.Write([]byte("ERROR 500"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	if err := tmpl.Execute(w, paste); err != nil {
		w.Write([]byte("ERROR 500"))
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func save(w http.ResponseWriter, r *http.Request) {
	p := db.Paste{}

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ERROR 500"))
	}
	r.Body.Close()

	if err := db.AddPaste(p); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"500"}`))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"id":"`+p.ID+`"}`))
}
