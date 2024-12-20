package posts

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4"
)

// Post struct with validation tags
type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title" validate:"required,min=1,max=255"`
	Content   string    `json:"content" validate:"required,min=1"`
	Category  string    `json:"category" validate:"required,min=1,max=100"`
	Tags      []string  `json:"tags" validate:"required,min=1,max=3"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Validate the post
func (p *Post) Validate() error {
	return validator.New().Struct(p)
}

// Create a post in the database
func Create(conn *pgx.Conn, post Post) (Post, error) {
	query := `
	INSERT INTO posts (title, content, category, tags, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, title, content, category, tags, created_at, updated_at`
	row := conn.QueryRow(context.Background(), query, post.Title, post.Content, post.Category, post.Tags, post.CreatedAt, post.UpdatedAt)
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.Category, &post.Tags, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return Post{}, err
	}
	return post, nil
}

// FindAll retrieves all posts from database
func FindAll(conn *pgx.Conn) ([]Post, error) {
	query := `SELECT * FROM posts`
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Category, &post.Tags, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// FindOne retrieves a single post by its ID
func FindOne(conn *pgx.Conn, id int) (Post, error) {
	query := `SELECT * FROM posts WHERE id = $1`
	row := conn.QueryRow(context.Background(), query, id)

	var post Post
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.Category, &post.Tags, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return Post{}, nil
	}
	return post, nil
}

// Update updates a post in the database by its ID
func Update(conn *pgx.Conn, post Post) (Post, error) {
	query := `
		UPDATE posts
		SET title = $1, content = $2, category = $3, tags = $4, updated_at = $5
		WHERE id = $6
		RETURNING id, title, content, category, tags, created_at, updated_at`
	row := conn.QueryRow(context.Background(), query, post.Title, post.Content, post.Category, post.Tags, post.UpdatedAt, post.ID)

	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.Category, &post.Tags, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return Post{}, err
	}
	return post, nil
}

// Remove removes a post from the database by its ID
func Remove(conn *pgx.Conn, id int) error {
	query := `DELETE FROM posts WHERE id = $1`
	_, err := conn.Exec(context.Background(), query, id)
	return err
}
