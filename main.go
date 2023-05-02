package main

import (
	"log"

	"github.com/AdityaRajSingh/go-blog-post/controllers"
	_ "github.com/AdityaRajSingh/go-blog-post/docs"
	"github.com/AdityaRajSingh/go-blog-post/initializers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

var API_URL string

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	initializers.InitDB(&config)
	API_URL = config.API_URL
}

// @title Blog-Post API
// @version 2.0
// @description This is a sample server
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 20.25.40.163:8000
// @BasePath /
// @schemes http
func main() {

	// Fiber instance
	app := fiber.New()
	micro := fiber.New()

	// Middleware
	app.Mount("/api", micro)
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))

	// Routes
	micro.Route("/blog-post", func(router fiber.Router) {
		router.Post("/", controllers.CreateBlogPost)
		router.Get("", controllers.FindAllBlogPosts)
	})
	micro.Route("/blog-post/:blogPostId", func(router fiber.Router) {
		router.Delete("", controllers.DeleteBlogPost)
		router.Get("", controllers.FindBlogPostById)
		router.Patch("", controllers.UpdateBlogPost)
	})
	micro.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Golang, Fiber, and GORM",
		})
	})

	micro.Get("/swagger/*", swagger.HandlerDefault)

	// Start Server
	log.Fatal(app.Listen(":8000"))
}
