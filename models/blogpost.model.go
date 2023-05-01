package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BlogPost struct {
	ID          string    `gorm:primary_key" json:"id,omitempty"`
	Title       string    `gorm:"varchar(255);not null" json:"title,omitempty"`
	Description string    `gorm:"type:text;not null" json:"description,omitempty"`
	Body        string    `gorm:"type:text;not null" json:"body,omitempty"`
	CreatedAt   time.Time `gorm:"not null" json:"createdAt,omitempty"`
	UpdatedAt   time.Time `gorm:"not null" json:"updatedAt,omitempty"`
}

func (blogPost *BlogPost) BeforeCreate(tx *gorm.DB) (err error) {
	blogPost.ID = uuid.NewString()
	return
}

var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

// swagger:parameters createBlogPostSchema
type CreateBlogPostSchema struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Body        string `json:"body" validate:"required"`
}

type UpdateBlogPostSchema struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Body        string `json:"body,omitempty"`
}
