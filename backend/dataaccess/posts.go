package dataaccess

import (
	"fmt"

	"github.com/owenyeo/sample-react-app/backend/database"
	"github.com/owenyeo/sample-react-app/backend/models"
)

func ListPosts(db *database.Database) ([]models.Post, error) {
	rows, err := db.Query("SELECT id, title, content, author, created_at FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author, &post.Date)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	fmt.Println(posts);
	return posts, nil
}

// AddPost adds a new post to the database
func AddPost(db *database.Database, newUser models.Post) error {

	_, err := db.Exec("INSERT INTO posts (title, content, author) VALUES ($1, $2, $3)", newUser.Title, newUser.Content, newUser.Author)

	if err != nil {
		return err
	}

	return nil
}
