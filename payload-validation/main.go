package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Username string `json:"username" validate:"required,max=10"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type SuccessResponse struct {
	Success bool `json:"success"`
	Data    User `json:"data"`
}

type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type CustomValidator struct {
	validator *validator.Validate
}

var validate = validator.New()

func (v *CustomValidator) Validate(data any) error {
	return v.validator.Struct(data)
}

func main() {
	registrationValidator := &CustomValidator{
		validator: validate,
	}
	app := fiber.New(fiber.Config{
		// Global custom error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if castedObject, ok := err.(validator.ValidationErrors); ok {
				for _, err := range castedObject {
					switch err.Tag() {
					case "required":
						return c.Status(fiber.StatusBadRequest).JSON(GlobalErrorHandlerResp{
							Success: false,
							Message: fmt.Sprintf("%s is required", err.Field()),
						})
					case "email":
						return c.Status(fiber.StatusBadRequest).JSON(GlobalErrorHandlerResp{
							Success: false,
							Message: fmt.Sprintf("%s is not valid email", err.Field()),
						})
					case "min":
						return c.Status(fiber.StatusBadRequest).JSON(GlobalErrorHandlerResp{
							Success: false,
							Message: fmt.Sprintf("%s  must not be less than 8 characters", err.Field()),
						})
					case "max":
						return c.Status(fiber.StatusBadRequest).JSON(GlobalErrorHandlerResp{
							Success: false,
							Message: fmt.Sprintf("%s must not be more than 10 characters", err.Field()),
						})
					}
				}
			}
			return c.Status(fiber.StatusInternalServerError).JSON(GlobalErrorHandlerResp{
				Success: false,
				Message: err.Error(),
			})
		},
	})

	app.Post("/register", func(c *fiber.Ctx) error {
		newUser := new(User)
		if err := c.BodyParser(newUser); err != nil {
			return err
		}

		// validation
		if err := registrationValidator.Validate(newUser); err != nil {
			return err
		}
		c.SendStatus(fiber.StatusCreated)
		return c.JSON(SuccessResponse{
			Success: true,
			Data:    *newUser,
		})
	})

	app.Listen(":3000")
}
