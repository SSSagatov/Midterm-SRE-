package db

import (
	"context"
	"errors"
	"strings"
	"time"

	"campus-connect/backend/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostRepository struct {
	pool *pgxpool.Pool
}

func NewPostRepository(pool *pgxpool.Pool) *PostRepository {
	return &PostRepository{pool: pool}
}

func (r *PostRepository) ListPosts(ctx context.Context) ([]models.Post, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT id, author, title, content, likes, created_at
		FROM posts
		ORDER BY created_at DESC
		LIMIT 50
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := make([]models.Post, 0)
	for rows.Next() {
		var p models.Post
		if err := rows.Scan(&p.ID, &p.Author, &p.Title, &p.Content, &p.Likes, &p.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, rows.Err()
}

func (r *PostRepository) CreatePost(ctx context.Context, post models.Post) (models.Post, error) {
	if strings.TrimSpace(post.Author) == "" || strings.TrimSpace(post.Title) == "" || strings.TrimSpace(post.Content) == "" {
		return models.Post{}, errors.New("author, title and content are required")
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := `
		INSERT INTO posts(author, title, content)
		VALUES($1, $2, $3)
		RETURNING id, author, title, content, likes, created_at
	`
	var created models.Post
	err := r.pool.QueryRow(ctx, query, post.Author, post.Title, post.Content).
		Scan(&created.ID, &created.Author, &created.Title, &created.Content, &created.Likes, &created.CreatedAt)
	return created, err
}

func (r *PostRepository) LikePost(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cmd, err := r.pool.Exec(ctx, `UPDATE posts SET likes = likes + 1 WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return errors.New("post not found")
	}
	return nil
}
