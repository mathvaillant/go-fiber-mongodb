package main

import (
	"example/go-fiber-mongodb/config"
	"example/go-fiber-mongodb/database"
	"example/go-fiber-mongodb/internal/user"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	// Init environment and database connection
	config, _ := config.NewConfig()
	mongoInstance, _ := database.NewDatabaseConnection(config.MongoURI, config.MongoDBName)

	// Create a new Fiber instance.
	app := fiber.New(fiber.Config{
		AppName:      "go-fiber-mongodb",
		ServerHeader: "Fiber",
	})

	// Use global middlewares.
	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(etag.New())
	app.Use(favicon.New())
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(limiter.New(limiter.Config{
		Max: 250,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
				"status":  "fail",
				"message": "You have requested too many in a single time-frame! Please wait another minute!",
			})
		},
	}))

	// Create repositories
	userRepository := user.NewUserRepository(mongoInstance)

	// Create services
	userService := user.NewUserService(userRepository)

	// Prepare endpoint handlers
	user.NewUserHandler(app.Group("/api/v1/users"), userService)

	// Prepare an endpoint for 'Not Found'.
	app.All("*", func(c *fiber.Ctx) error {
		errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())

		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"status":  "fail",
			"message": errorMessage,
		})
	})

	log.Fatal(app.Listen(":" + config.Port))
}
