package models

import "time"

type Post struct {
	ID        int64     `json:"id"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Likes     int       `json:"likes"`
	CreatedAt time.Time `json:"created_at"`
}
