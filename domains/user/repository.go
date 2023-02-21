package user

import (
	"app/pkg/storages/pg"
	"github.com/gofiber/fiber/v2"
)

func GetAllUsersRepository(ctx *fiber.Ctx) error {
	return ctx.SendString("Get All Users")
}

func GetOneUserRepository(ctx *fiber.Ctx) error {
	return ctx.SendString("Get One User")
}

func CreateUserRepository(user *User) {
	var db = pg.Connect().Conn
	db.Create(&user)
}

func UpdateUserRepository(ctx *fiber.Ctx) error {
	return ctx.SendString("Update User")
}

func DeleteUserRepository(ctx *fiber.Ctx) error {
	return ctx.SendString("Delete User")
}
