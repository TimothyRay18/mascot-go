package controller

import (
	"mascot/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) AddNewUser(context *fiber.Ctx) error {
	userApp := models.UserApp{}

	err := context.BodyParser(&userApp)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	existingUser := models.UserApp{}
	r.DB.Where("username = ?", *userApp.Username).First(&existingUser)
	if existingUser.Username != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "user has been taken"})
		return err
	}
	defaultLevel := 0
	userApp.Level = &defaultLevel

	err = r.DB.Create(&userApp).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not add user"})
		return err
	}

	data := map[string]interface{}{
		"user":   userApp.Username,
		"access": userApp.ViewAccess,
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "user has been added",
		"data":    data})
	return nil
}

func (r *UserRepository) Login(context *fiber.Ctx) error {
	userApp := models.UserApp{}

	err := context.BodyParser(&userApp)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	cekUsername := GetUsernameFromToken(context)
	user := models.UserApp{}
	if cekUsername == "" {
		// If user credentials is correct | continue
		if err := r.DB.Where("username = ? AND password = ?", *userApp.Username, *userApp.Password).First(&user).Error; err == nil {
			GenerateToken(context, *user.Username, 0)
			context.Status(http.StatusOK).JSON(
				&fiber.Map{"message": "login success"})
			return nil
		} else {
			context.Status(http.StatusBadRequest).JSON(
				&fiber.Map{"message": "user not found"})
			return nil
		}
	} else {
		context.Status(http.StatusNotAcceptable).JSON(&fiber.Map{
			"message": "already logged in as " + *userApp.Username})
		return nil
	}
}

func (r *UserRepository) Logout(context *fiber.Ctx) error {
	ResetUserToken(context)

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "logout success"})

	return nil
}
