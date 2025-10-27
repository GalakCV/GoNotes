package models

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Note struct {
	Id pgtype.Numeric        
	Title pgtype.Text
	Content pgtype.Text 
	Color pgtype.Text     
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}