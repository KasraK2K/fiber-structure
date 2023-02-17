package users

import (
	"app/databases/pg"
	"app/models"
	"encoding/json"
	"fmt"
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
		return err
	}

	//Create User
	result := db.Create(&user)
	fmt.Println(result)

	return c.SendString("Create User")
}

func UpdateUser(c *fiber.Ctx) error {
	return c.SendString("Update User")
}

func DeleteUser(c *fiber.Ctx) error {
	return c.SendString("Delete User")
}
