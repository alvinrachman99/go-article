package repository

import (
	"article/models"
	"database/sql"
	"time"
)

type PostsRepository struct {
	DB *sql.DB
}

func NewPostsRepository(db *sql.DB) *PostsRepository {
	return &PostsRepository{DB: db}
}

func (r *PostsRepository) Create(posts *models.Posts) error {
	query := "INSERT INTO posts (title, content, category, status) VALUES (?, ?, ?, ?)"
	_, err := r.DB.Exec(query, posts.Title, posts.Content, posts.Category, posts.Status)

	return err
}

func (r *PostsRepository) GetPosts(limit int, offset int) ([]models.Posts, error) {
	query := "SELECT id, title, content, category, status, created_at, updated_at FROM posts LIMIT ? OFFSET ?"
	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var createdAtStr, updatedAtStr string

	var posts []models.Posts
	for rows.Next() {
		var post models.Posts
		if err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.Category, &post.Status, &createdAtStr, &updatedAtStr); err != nil {
			return nil, err
		}

		post.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAtStr)
		post.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAtStr)

		posts = append(posts, post)
	}

	return posts, nil
}

func (r *PostsRepository) GetPostsById(id int) (*models.Posts, error) {
	query := "SELECT id, title, content, category, status, created_at, updated_at FROM posts WHERE id = ?"
	row := r.DB.QueryRow(query, id)

	var createdAtStr, updatedAtStr string

	var post models.Posts
	if err := row.Scan(&post.Id, &post.Title, &post.Content, &post.Category, &post.Status, &createdAtStr, &updatedAtStr); err != nil {
		return nil, err
	}

	post.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAtStr)
	post.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAtStr)

	return &post, nil
}

func (r *PostsRepository) Update(post *models.Posts) error {
	query := "UPDATE posts SET title = ?, content = ?, category = ?, status = ? WHERE id = ?"
	_, err := r.DB.Exec(query, post.Title, post.Content, post.Category, post.Status, post.Id)

	return err
}

func (r *PostsRepository) Delete(id int) error {
	query := "DELETE FROM posts WHERE id = ?"
	_, err := r.DB.Exec(query, id)

	return err
}

func (r *PostsRepository) UpdateStatusTrash(id int) error {
	query := "UPDATE posts SET status = 'trash' WHERE id = ? "
	_, err := r.DB.Exec(query, id)

	return err
}
