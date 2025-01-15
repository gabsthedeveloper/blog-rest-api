package models

import (
	"time"

	"example.com/blog-rest-api/db"
)

type Post struct {
	ID          int64
	Title       string    `binding:"required"`
	Description string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func GetAllPosts() ([]Post, error) {
	query := "SELECT * FROM posts"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []Post

	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Description, &post.DateTime, &post.UserID)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func GetPostByID(id int64) (*Post, error) {
	query := "SELECT * FROM posts WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var post Post
	err := row.Scan(&post.ID, &post.Title, &post.Description, &post.DateTime, &post.UserID)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (post *Post) Save() error {
	query := `
	INSERT INTO posts(title, description, dateTime, user_id)
	VALUES (?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(post.Title, post.Description, post.DateTime, post.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	post.ID = id
	return err
}

func (post Post) Update() error {
	query := `
	UPDATE posts
	SET title = ?, description = ?, dateTime = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(post.Title, post.Description, post.DateTime, post.ID)
	return err
}

func (post Post) Delete() error {
	query := "DELETE FROM posts WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(post.ID)
	return err
}
