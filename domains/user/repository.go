package user

import (
	"app/databases/pg"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func migrate() (*gorm.DB, error) {
	var db = pg.Connect().Conn

	//Create Table If Not Exist
	err := db.AutoMigrate(&User{})
	return db, err
}

func GetAllUsersRepository(ctx *fiber.Ctx) error {
	return ctx.SendString("Get All Users")
}

func GetOneUserRepository(ctx *fiber.Ctx) error {
	return ctx.SendString("Get One User")
}

func CreateUserRepository(user *User) (bool, error) {
	var db, migrateError = migrate()
	if migrateError != nil {
		return false, migrateError
	} else {
		db.Create(&user)
		return true, nil
	}
}

func UpdateUserRepository(ctx *fiber.Ctx) error {
	return ctx.SendString("Update User")
}

func DeleteUserRepository(ctx *fiber.Ctx) error {
	return ctx.SendString("Delete User")
}
