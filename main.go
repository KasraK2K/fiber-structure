package main

import (
	"app/databases/pg"
	"app/models"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	app := fiber.New()

	//Middlewares
	app.Use(cache.New())
	app.Use(compress.New())
	app.Use(cors.New())
	app.Use(csrf.New())
	app.Use(etag.New())
	app.Use(favicon.New())
	app.Use(limiter.New(limiter.Config{Max: 100, Expiration: 60 * time.Second}))
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	//Configs
	app.Static("/", "./public")

	//Routes
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Default Metrics Page"}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	//Fill User Struct
	var user = models.User{
		FirstName:  "Kasra",
		Surname:    "Karami",
		Email:      "Kasra_K2K@yahoo.com",
		Password:   "12345678",
		Permission: "1111111111",
		IsActive:   true,
		IsAdmin:    true,
		Phone:      "09183619290",
	}
	//Validate User Struct
	validationError := user.Validate()
	if validationError.Errors != nil {
		data, _ := json.Marshal(validationError.Errors)
		fmt.Println(string(data))
	}
	//Create Table If Not Exist
	var db = pg.Connect().Conn
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
	//Create User
	result := db.Create(&user)
	fmt.Println(result)

	//Init & Log
	//log.Fatal(app.Listen(":3000"))
}
