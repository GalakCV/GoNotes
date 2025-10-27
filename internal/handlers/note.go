package handlers

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"udemy.com/galakcv/aulago/internal/handlers/apperror"
	"udemy.com/galakcv/aulago/internal/repositories"
)

type noteHandler struct {
	repo repositories.NoteRepository
}

func NewNoteHandler(repo repositories.NoteRepository) *noteHandler {
	return &noteHandler{repo: repo}
}

func (nh *noteHandler) NoteList(w http.ResponseWriter, r *http.Request){

	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}
	files := []string{
		"views/templates/base.html",
		"views/templates/pages/home.html",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "ERROR", http.StatusInternalServerError)
		return
	}

	notes, err := nh.repo.List(r.Context())
	if err != nil {
		return 
	}
	
	t.ExecuteTemplate(w, "base", newNoteResponseFromNoteList(notes))
}

func (nh *noteHandler) NoteView(w http.ResponseWriter, r *http.Request) error {
	idParam := r.URL.Query().Get("id") 
	if idParam == ""{ 
		return apperror.WithStatus(errors.New("note is obrigatory"), http.StatusBadRequest)
	}

	id,err := strconv.Atoi(idParam)
	if err != nil{
		return err
	}

	files := []string{
		"views/templates/base.html",
		"views/templates/pages/noteView.html",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		return errors.New("error found")
	}

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()
	if id == 1{
		cancel()
	}
	note, err := nh.repo.GetById(ctx, id)
	if err != nil {
		return err
	}
	
	buff := &bytes.Buffer{}
	err = t.ExecuteTemplate(buff, "base", newNoteResponseFromNote(note))
	if err != nil {
		return err 
	}

	buff.WriteTo(w)
	return nil
}

func (nh *noteHandler) NoteNew(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"views/templates/base.html",
		"views/templates/pages/noteNew.html",
	}
	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "ERROR", http.StatusInternalServerError)
		return
	}
	
	t.ExecuteTemplate(w, "base", nil)
}


func (nh *noteHandler) NoteCreate(w http.ResponseWriter, r *http.Request){
	
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Invalid request", http.StatusMethodNotAllowed)
		return 
	}
	fmt.Fprintf(w, "Creating a new note...")
}
