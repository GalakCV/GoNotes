package handlers

import "udemy.com/galakcv/aulago/internal/models"

type NoteResponse struct {
	Id      int
	Title   string
	Content string
}

func newNoteResponseFromNote(note *models.Note)(res NoteResponse){
	res.Id = int(note.Id.Int.Int64())
	res.Title = note.Title.String
	res.Content = note.Title.String
	return
}

func newNoteResponseFromNoteList(notes []models.Note)(res []NoteResponse){
	for _, note := range notes{
		res = append(res, newNoteResponseFromNote(&note))
	}
	return
}