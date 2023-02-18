package main

import (
	"app/handlers"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
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
	app := fiber.New()

	//Middlewares
	app.Use(cache.New())
	app.Use(compress.New())
	app.Use(cors.New())
	//app.Use(csrf.New())
	app.Use(etag.New())
	app.Use(favicon.New())
	app.Use(limiter.New(limiter.Config{Max: 100, Expiration: 60 * time.Second}))
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	//Configs
	app.Static("/", "./public")

	//Routes
	handlers.Routes(app)

	//Init & Log
	log.Fatal(app.Listen(":3000"))
}
