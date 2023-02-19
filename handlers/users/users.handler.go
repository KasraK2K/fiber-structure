package users

import (
	"app/databases/pg"
	"app/models"
	"github.com/gofiber/fiber/v2"
)

func Routes(router fiber.Router) {
	userGroup := router.Group("/users")

	userGroup.Get("/", GetAllUsers)
	userGroup.Get("/:id", GetOneUser)
	userGroup.Post("/", CreateUser)
	userGroup.Patch("/:id", UpdateUser)
	userGroup.Get("/:id", DeleteUser)
}

func GetAllUsers(c *fiber.Ctx) error {
	return c.SendString("Get All Users")
}

func GetOneUser(c *fiber.Ctx) error {
	return c.SendString("Get One User")
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	parseError := c.BodyParser(user)
	if parseError != nil {
		return parseError
	}

	////Validate User Struct
	validationError := user.Validate()
	if validationError.Errors != nil {
		return c.JSON(validationError)
	}

	//Create Table If Not Exist
	var db = pg.Connect().Conn
	migrateError := db.AutoMigrate(&models.User{})
	if migrateError != nil {
		return c.JSON(migrateError)
	}

	//Create User & send response
	db.Create(&user)
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	return c.SendString("Update User")
}

func DeleteUser(c *fiber.Ctx) error {
	return c.SendString("Delete User")
}
