package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"github.com/jackc/pgx/v5/pgxpool"
	"udemy.com/galakcv/aulago/internal/handlers"
	"udemy.com/galakcv/aulago/internal/repositories"
)

func main() {
	slog.SetDefault(LoggerNovo(os.Stderr, slog.LevelInfo))

	dbpool, err := pgxpool.New(context.Background(), "postgres://postgres:galak@localhost:5432/postgres")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	slog.Info("Conection sucessful")
	defer dbpool.Close()
	fmt.Printf("Iniciando servidor Go") 
	
	Mux := http.NewServeMux()
	staticHandler := http.FileServer(http.Dir("views/static"))
	Mux.Handle("/static/", http.StripPrefix("/static/", staticHandler))

	noteRepo := repositories.NewNoteRepository(dbpool)

	noteHandler := handlers.NewNoteHandler(noteRepo)

	Mux.HandleFunc("/", noteHandler.NoteList)
	Mux.Handle("/note/view", handlers.HandlerWithError(noteHandler.NoteView))
	Mux.HandleFunc("/note/create", noteHandler.NoteCreate)
	Mux.HandleFunc("/note/new", noteHandler.NoteNew)

	http.ListenAndServe(":5000", Mux)

}