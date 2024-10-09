package posts

import "time"

// Post is the struct of post
type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Category  string    `json:"category"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var posts = []Post{
	{
		ID:        1,
		Title:     "Test Post",
		Content:   "This is a test post",
		Category:  "Test Category",
		Tags:      []string{"Test Tag", "Test Tag 2"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        2,
		Title:     "Test Post 2",
		Content:   "This is a test post 2",
		Category:  "Test Category 2",
		Tags:      []string{"Test Tag 3", "Test Tag 4"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        3,
		Title:     "Test Post 3",
		Content:   "This is a test post 3",
		Category:  "Test Category 3",
		Tags:      []string{"Test Tag 5", "Test Tag 6"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

// FindAll returns a list of posts
func FindAll() ([]Post, error) {
	return posts, nil
}

// FindOne returns a single post by id
func FindOne(id int) (Post, error) {
	for _, post := range posts {
		if post.ID == id {
			return post, nil
		}
	}
	return Post{}, nil
}

// Create creates a post
func Create() (Post, error) {
	return Post{
		ID:        len(posts) + 1,
		Title:     "My first post",
		Content:   "Such a nice app",
		Category:  "Social",
		Tags:      []string{"social", "lifestyle"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
