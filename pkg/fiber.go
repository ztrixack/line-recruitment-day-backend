package pkg

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"election-service/configs"
	"election-service/pkg/ws"
)

func InitRest() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork: configs.Server.Prefork,
		AppName: configs.App.Name,
	})

	app.Use(requestid.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "OPTIONS,GET,POST",
		AllowHeaders:     "Origin, X-Requested-With, Content-Type, Accept, Authorization",
		AllowCredentials: true,
		MaxAge:           43200000000000,
	}))
	app.Use(logger.New(logger.Config{
		Format:   "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
		TimeZone: configs.App.Timezone,
	}))

	return app
}

func InitWsHub() *ws.Hub {
	hub := ws.Hub{
		Election: make(map[string]string),
		Vote:     make(map[string]string),
	}
	return &hub
}
