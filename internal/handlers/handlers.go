package handlers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/NoistNT/blogging-platform-api/internal/posts"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

// CreatePost Handler
func CreatePost(c *gin.Context, conn *pgx.Conn) {
	var post posts.Post
	if err := c.BindJSON(&post); err != nil {
		log.Printf("Error binding JSON in CreatePost: %v", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	createdPost, err := posts.Create(conn, post)
	if err != nil {
		log.Printf("Error creating post in database: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to create post"})
		return
	}
	c.IndentedJSON(http.StatusCreated, createdPost)
}

// GetPosts Handler
func GetPosts(c *gin.Context, conn *pgx.Conn) {
	posts, err := posts.FindAll(conn)
	if err != nil {
		log.Printf("Error fetching posts from database: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to get posts"})
		return
	}
	c.IndentedJSON(http.StatusOK, posts)
}

// GetPost Handler
func GetPost(c *gin.Context, conn *pgx.Conn) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Error parsing post ID: %v", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
		return
	}

	post, err := posts.FindOne(conn, id)
	if err != nil {
		log.Printf("Error fetching post with ID %d: %v", id, err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to retrieve post"})
		return
	}

	if post.ID == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, post)
}

// RemovePost Handler
func RemovePost(c *gin.Context, conn *pgx.Conn) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Invalid post ID: %d: %v", id, err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
		return
	}

	err = posts.Remove(conn, id)
	if err != nil {
		log.Printf("Error removing post with ID %d: %v", id, err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to remove post"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "post removed successfully"})
}
