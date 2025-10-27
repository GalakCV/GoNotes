package repositories

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"udemy.com/galakcv/aulago/internal/models"
)

type NoteRepository interface {
	List(ctx context.Context) ([]models.Note, error)
	GetById(ctx context.Context, id int)(*models.Note, error)
	Create(ctx context.Context, title, content, color string) (*models.Note, error)
	Update(ctx context.Context, id int, title, content, color string) (*models.Note, error)
	Delete(ctx context.Context, id int) (*models.Note, error)
}


func NewNoteRepository(dbpool *pgxpool.Pool) NoteRepository{
	return &noteRepository{db: dbpool}
}

type noteRepository struct {
	db *pgxpool.Pool
}

func (nr *noteRepository) List(ctx context.Context) ([]models.Note, error) {
	var list []models.Note
	rows, err := nr.db.Query(ctx, "select id, title, content, color, created_at, updated_at from notes")

	if err != nil {
		return list,err
	}
	defer rows.Close()
	for rows.Next() {
		var note models.Note
		err := rows.Scan(&note.Id, &note.Title, &note.Content, &note.Color, &note.CreatedAt, &note.UpdatedAt)
		if err != nil {
			return list, err
		}
		list = append(list, note)
	}
	return list, nil
}

func (nr *noteRepository) GetById(ctx context.Context, id int)(*models.Note, error) {
	var note models.Note
	noteGetQuerry := `select id, title, content, color, created_at, updated_at from notes where id = $1`
	row := nr.db.QueryRow(ctx, noteGetQuerry, id)
	if err := row.Scan(&note.Id, &note.Title, &note.Content, &note.Color, &note.CreatedAt, &note.UpdatedAt); err != nil {
		return &note, err 
	}
	return &note, nil
}

func (nr *noteRepository) Create(ctx context.Context, title, content, color string) (*models.Note, error) {
	var note models.Note

	note.Title = pgtype.Text{String: title, Valid: true}
	note.Content = pgtype.Text{String: content, Valid: true}
	note.Color = pgtype.Text{String: color, Valid: true}

	CreateQuerry := `INSERT INTO notes (title, content, color) VALUES ($1, $2, $3)
	RETURNING id, created_at
	`
	row := nr.db.QueryRow(ctx, CreateQuerry, title, content, color)
	if err := row.Scan(&note.Id, &note.CreatedAt); err != nil {
		return &note, err
	}
	return &note, nil
}


func (nr *noteRepository) Update(ctx context.Context, id int, title, content, color string) (*models.Note, error) {
	var note models.Note

	query := `UPDATE notes set title = $1, content= $2, color = $3 WHERE id = $4`

	if len(title) > 0 {
		note.Title = pgtype.Text{String: title, Valid: true}
	}
	
	if len(content) > 0 {
		note.Content = pgtype.Text{String: content, Valid: true}
	}
	
	if len(color) > 0 {
		note.Color = pgtype.Text{String: color, Valid: true}
	}
	_, err := nr.db.Exec(ctx, query, title, content, color, id)
	if err != nil {
		return &note, nil
	}
	return &note, nil
}

func (nr *noteRepository) Delete(ctx context.Context, id int) (*models.Note, error) {
	var note models.Note

	DeleteQuery := `DELETE from notes where id = $1`
	_, err := nr.db.Exec(ctx, DeleteQuery, id)
	if err != nil {
		return nil, err
	}
	return &note, nil
}