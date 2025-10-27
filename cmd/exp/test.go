package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

func main() {

	var err error
	db_url := "postgres://postgres:galak@localhost:5432/postgres"
	conn, err = pgx.Connect(context.Background(), db_url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection sucessful !")
	defer conn.Close(context.Background())
	selectAll()
}

type Post struct {
	id int
	title string
	content string
	author string
}

func selectAll(){
	query := "SELECT id, title, content, author FROM posts"
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	if rows.Err() != nil {
		panic(rows.Err())
	}
	var posts []Post
	for rows.Next(){
		var post Post
		err = rows.Scan(&post.id, &post.title, &post.content, &post.author)
		if err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}
	for _, post := range posts {
		fmt.Println(post)
	}
}