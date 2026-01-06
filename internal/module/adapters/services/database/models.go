package database

import (
	"time"

	"github.com/google/uuid"
)

// Base model para evitar repetir el campo ID UUID
type Base struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
}

type User struct {
	Base
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	// Relaciones
	Profile Profile `gorm:"foreignKey:UserID"`
	Posts   []Post  `gorm:"foreignKey:UserID"`
}

type Profile struct {
	Base
	UserID    uuid.UUID `gorm:"type:uuid;unique;not null"`
	Alias     string    `gorm:"unique;not null"`
	FirstName string    `gorm:"type:text;not null"`
	LastName  string    `gorm:"type:text;not null"`

	BirthDate string `gorm:"type:date"`
}

type Post struct {
	Base
	UserID    uuid.UUID `gorm:"type:uuid;index;not null"`
	Message   string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	// Relación para traer los likes del post
	Likes []Like `gorm:"foreignKey:PostID"`
}

type Like struct {
	Base
	PostID    uuid.UUID `gorm:"type:uuid;index;not null"`
	UserID    uuid.UUID `gorm:"type:uuid;index;not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
