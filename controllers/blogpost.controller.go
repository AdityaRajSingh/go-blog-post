package controllers

import (
	"strings"
	"time"

	"github.com/AdityaRajSingh/go-blog-post/initializers"
	"github.com/AdityaRajSingh/go-blog-post/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// CreateBlogPost func for creates a new blogPost.
// @Description Create a new blogPost.
// @Summary create a new blogPost
// @TagsBlogPosts
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param content body string true "Content"
// @Param category body string true "Category"
// @Success 200 {object} models.BlogPost
// @Router /api/blog-post [post]
func CreateBlogPost(c *fiber.Ctx) error {
	var payload *models.CreateBlogPostSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	now := time.Now()
	newBlogPost := models.BlogPost{
		Title:       payload.Title,
		Description: payload.Description,
		Body:        payload.Body,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	result := initializers.DB.Create(&newBlogPost)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Title already exist, please use another title"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"blogPost": newBlogPost}})
}

// FindAllBlogPosts func gets all exists blogPosts.
// @Description Get all exists blogPosts.
// @Summary get all exists blogPosts
// @TagsBlogPosts
// @Accept json
// @Produce json
// @Success 200 {array} models.BlogPost
// @Router /api/blog-post [get]
func FindAllBlogPosts(c *fiber.Ctx) error {
	var blogPosts []models.BlogPost
	results := initializers.DB.Find(&blogPosts)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(blogPosts), "blogPosts": blogPosts})
}

// UpdateBlogPost func for updates blogPost by given ID.
// @Description Update blogPost.
// @Summary update blogPost
// @TagsBlogPosts
// @Accept json
// @Produce json
// @Param id body string true "BlogPost ID"
// @Param content body string true "Content"
// @Param category body string true "Category"
// @Success 201 {string} status "ok"
// @Router /api/blog-post [patch]
func UpdateBlogPost(c *fiber.Ctx) error {
	blogPostId := c.Params("blogPostId")

	var payload *models.UpdateBlogPostSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	var blogPost models.BlogPost
	result := initializers.DB.First(&blogPost, "id = ?", blogPostId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No blogPost with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	updates := make(map[string]interface{})
	if payload.Title != "" {
		updates["title"] = payload.Title
	}
	if payload.Body != "" {
		updates["category"] = payload.Body
	}
	if payload.Description != "" {
		updates["content"] = payload.Description
	}

	updates["updated_at"] = time.Now()

	initializers.DB.Model(&blogPost).Updates(updates)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"blogPost": blogPost}})
}

// FindBlogPostById func gets blogPosts by given ID or 404 error.
// @Description Get blogPosts by given ID.
// @Summary get blogPosts by given ID
// @TagsBlogPosts
// @Accept json
// @Produce json
// @Param id path string true "blogPostId"
// @Success 200 {object} models.BlogPost
// @Router /api/blog-post/{id} [get]
func FindBlogPostById(c *fiber.Ctx) error {
	blogPostId := c.Params("blogPostId")

	var blogPost models.BlogPost
	result := initializers.DB.First(&blogPost, "id = ?", blogPostId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No blogPost with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"blogPost": blogPost}})
}

// DeleteBlogPost func for deletes book by given ID or 404 error.
// @Description Delete blogPost by given ID.
// @Summary delete blogPost by given ID
// @TagsBlogPosts
// @Accept json
// @Produce json
// @Param id path string true "blogPostId"
// @Success 200 {object} models.BlogPost
// @Router /api/blog-post/{id} [delete]
func DeleteBlogPost(c *fiber.Ctx) error {
	blogPostId := c.Params("blogPostId")

	result := initializers.DB.Delete(&models.BlogPost{}, "id = ?", blogPostId)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No blogPost with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
