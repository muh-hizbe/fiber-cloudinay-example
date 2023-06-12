package main

import (
	"fiber-cloudinary/database"
	"fiber-cloudinary/model"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	urlCloudinary := "cloudinary://ap-key:secret-key@cloud-name"
	
	database.InitDatabase()
	database.AutoMigration()
	app := fiber.New()

	app.Post("/user", func(c *fiber.Ctx) error {
		var user model.User

		_ = c.BodyParser(&user)

		fileHeader, _ := c.FormFile("image")
		log.Println(fileHeader.Filename)

		return c.JSON(fiber.Map{
			"data": user,
		})
	})

	app.Listen(":8000")
	fmt.Println("server running up on port 8000")
}
