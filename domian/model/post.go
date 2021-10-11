package model

import "time"

type Post struct {
	ID int64 `grom:"PRIMARY_KEY AUTO_INCREMENT" `
	Content string `grom:"not null"`
	Title string `grom:"not null"`
	UserID int64 `grom:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
