package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/AkulinIvan/ToDo-crud/internal/api/middleware"
	"github.com/AkulinIvan/ToDo-crud/internal/service"
)

type Routers struct {
	Service service.Service
}

func NewRouters(r *Routers, token string) *fiber.App {
	app := fiber.New()

	// Настройка CORS (разрешенные методы, заголовки, авторизация)
	app.Use(cors.New(cors.Config{
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowHeaders:     "Accept, Authorization, Content-Type, X-CSRF-Token, X-REQUEST-SomeID",
		ExposeHeaders:    "Link",
		AllowCredentials: false,
		MaxAge:           300,
	}))
	apiGroup := app.Group("/v1/", middleware.Authorization(token))

	apiGroup.Post("task/", r.Service.CreateTask)
	apiGroup.Get("tasks/", r.Service.GetTasks)
	apiGroup.Get("task/:id", r.Service.GetTask)
	apiGroup.Put("task/:id", r.Service.UpdateTask)
	apiGroup.Delete("task/:id", r.Service.DeleteTask)
	
	return app
}
