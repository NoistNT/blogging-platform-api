package main

import (
	"context"
	"log"

	"github.com/NoistNT/blogging-platform-api/cmd/migrate"
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

	// Connect to the database
	conn, err := pgx.Connect(context.Background(), cfg.DBURL)
	if err != nil {
		log.Fatalf("unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	// Check connection
	if err := conn.Ping(context.Background()); err != nil {
		log.Fatalf("unable to ping database: %v\n", err)
	}

	// Run migration
	if err := migrate.Migrate(conn); err != nil {
		log.Fatalf("failed to migrate database: %v\n", err)
	}

	router.POST("/posts", func(c *gin.Context) {
		handlers.CreatePost(c, conn)
	})
	router.GET("/posts", func(c *gin.Context) {
		handlers.GetPosts(c, conn)
	})
	router.GET("/posts/:id", func(c *gin.Context) {
		handlers.GetPost(c, conn)
	})
	router.DELETE("/posts/:id", func(c *gin.Context) {
		handlers.RemovePost(c, conn)
	})

	log.Fatalf("failed to start server: %v", router.Run(":8080"))
}
