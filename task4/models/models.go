package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Posts    []Post
	Comments []Comment
}

type Post struct {
	gorm.Model
	Title     string `gorm:"not null"`
	Content   string `gorm:"type:text;not null"`
	UserID    uint   `gorm:"not null"`
	User      User   `gorm:"foreignKey:UserID"`
	Comments  []Comment
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Comment struct {
	gorm.Model
	Content   string `gorm:"type:text;not null"`
	UserID    uint   `gorm:"not null"`
	PostID    uint   `gorm:"not null"`
	User      User   `gorm:"foreignKey:UserID"`
	Post      Post   `gorm:"foreignKey:PostID"`
	CreatedAt time.Time
}
