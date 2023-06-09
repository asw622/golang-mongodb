package routes

import (
	"golang-mongodb/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App){
	app.Post("/user", controllers.CreateUser)
	app.Get("/user/:userId", controllers.GetUser)
	app.Put("/user/:userId", controllers.EditUser)
	app.Delete("/user/:userId", controllers.DeleteUser)
	app.Get("/users", controllers.GetListUsers)
}