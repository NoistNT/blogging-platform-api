package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/NoistNT/blogging-platform-api/internal/posts"
	"github.com/gin-gonic/gin"
)

// GetPosts Handler
func GetPosts(c *gin.Context) {
	posts, err := posts.FindAll()
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, posts)
}

// GetPost Handler
func GetPost(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err)
	}

	post, err := posts.FindOne(id)
	if err != nil {
		log.Fatal(err)
	}

	if post.ID == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, post)
}

// CreatePost Handler
func CreatePost(c *gin.Context) {
	post, _ := posts.Create()
	c.IndentedJSON(http.StatusCreated, post)
}
