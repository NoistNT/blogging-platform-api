package main

import (
	"context"
	"log"

	"github.com/NoistNT/blogging-platform-api/config"
	"github.com/NoistNT/blogging-platform-api/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func main() {
	router := gin.Default()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the databse
	conn, err := pgx.Connect(context.Background(), cfg.DBURL)
	if err != nil {
		log.Fatalf("unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	// Check connection
	if err := conn.Ping(context.Background()); err != nil {
		log.Fatalf("unable to ping database: %v\n", err)
	}

	router.POST("/posts", handlers.CreatePost)
	router.GET("/posts", handlers.GetPosts)
	router.GET("/posts/:id", handlers.GetPost)

	log.Fatalf("failed to start server: %v", router.Run(":8080"))
}
