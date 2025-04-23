package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

/*
id         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY
name       VARCHAR(255) NOT NULL
tel        VARCHAR(255)
email      VARCHAR(255)
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
*/

// User represents the user model
type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id" example:"1"`              // User ID
	Name      string    `gorm:"not null" json:"name" validate:"required" example:"John Doe"` // User name
	Tel       *string   `json:"tel" validate:"omitempty,len=11" example:"01012345678"`       // Phone number
	Email     *string   `json:"email" validate:"omitempty,email" example:"john@example.com"` // Email address
	CreatedAt time.Time `json:"-" copier:"-"`                                                // Created timestamp (hidden)
	UpdatedAt time.Time `json:"-" copier:"-"`                                                // Updated timestamp (hidden)
}

// Validate runs validation on the User model
func (u *User) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(u); err != nil {
		return err.(validator.ValidationErrors)
	}
	return nil
}
