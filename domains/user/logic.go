package user

import "github.com/gofiber/fiber/v2"

func GetAllUsersLogic(ctx *fiber.Ctx) error {
	return ctx.SendString("Get All Users")
}

func GetOneUserLogic(ctx *fiber.Ctx) error {
	return ctx.SendString("Get One User")
}

func CreateUserLogic(ctx *fiber.Ctx) error {
	user := new(User)
	parseError := ctx.BodyParser(user)
	if parseError != nil {
		return parseError
	}

	//Validate User Struct
	validationError := user.Validate()
	if validationError.Errors != nil {
		return ctx.JSON(validationError)
	}

	result, err := CreateUserRepository(user)
	if result != true {
		return ctx.JSON(err)
	} else {
		return ctx.JSON(user)
	}
}

func UpdateUserLogic(ctx *fiber.Ctx) error {
	return ctx.SendString("Update User")
}

func DeleteUserLogic(ctx *fiber.Ctx) error {
	return ctx.SendString("Delete User")
}
