package model

import (
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type User struct {
	ID        string    `gorm:"primaryKey;size:26"`
	FirstName string    `gorm:"type:varchar(255);not null"`
	LastName  string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"type:varchar(255);unique;not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"type:timestamp(3);not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamp(3);not null"`
}

// Hook tự động tạo UUID khi insert
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		t := time.Now().UTC()
		entropy := rand.New(rand.NewSource(t.UnixNano()))
		u.ID = ulid.MustNew(ulid.Timestamp(t), entropy).String()
	}
	return nil
}
