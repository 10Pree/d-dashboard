package controller

import (
	"d-dashboard/connect"
	"d-dashboard/models"
	"errors"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"passwors"`
	Email    string `json:"email"`
}

type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreatedResponseUser(userModel models.User) User {
	return User{ID: userModel.ID, Username: userModel.Username, Email: userModel.Email, Password: userModel.Password}
}
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CreatedUser(c *fiber.Ctx) error {
	var userRequest UserRequest

	if err := c.BodyParser(&userRequest); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	hashedPassword, err := hashPassword(userRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	user := models.User{
		Username: userRequest.Username,
		Email:    userRequest.Email,
		Password: hashedPassword,
	}

	connect.Database.Db.Create(&user)
	responseUser := CreatedResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	connect.Database.Db.Find(&users)
	responseUsers := []User{}
	for _, user := range users {
		responseUser := CreatedResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}
	return c.Status(200).JSON(responseUsers)
}

func findUser(id int, user *models.User) error {
	connect.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please id is an integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreatedResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please id is an integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpateUser struct {
		Username string `json:"username"`
		Password string `json:"-"`
		Email    string `json:"email"`
	}

	var updateData UpateUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.Username = updateData.Username
	user.Password = updateData.Password
	user.Email = updateData.Email

	connect.Database.Db.Save(&user)

	responseUser := CreatedResponseUser(user)

	return c.Status(500).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please id is an integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := connect.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully Deleted User")

}
