package migrate

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// Migrate creates the posts table in the database if it doesn't exist
func Migrate(conn *pgx.Conn) error {
	query := `
	CREATE TABLE IF NOT EXISTS posts (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		content TEXT NOT NULL,
		category VARCHAR(100),
		tags TEXT[],
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := conn.Exec(context.Background(), query)
	return err
}
