package main

import (
	"authn-service-demo/api/middlewares"
	"authn-service-demo/api/routes"
	"authn-service-demo/infrastructure/config"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

const (
	PATH = "./config.json"
)

func main() {
	// Extract configuration params
	//  config.ExtractConfigParams(PATH, &config.GlobalConfigParams)
	config.LoadConfigParamsFromEnv(&config.GlobalConfigParams)
	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName:      "My Authn Service",
		ServerHeader: "Fiber",
	})
	// Add middlewares and routes
	middlewares.InitFiberMiddlewares(app, routes.InitPublicRoutes, routes.InitProtectedRoute)

	fmt.Println("Hello their this is my http server")

	// Server listens
	fmt.Println("Server listen on port:", config.GlobalConfigParams.ListenPort)
	err := app.Listen(fmt.Sprintf("%v:%v",
		config.GlobalConfigParams.ListenIP,
		config.GlobalConfigParams.ListenPort))
	log.Fatal(err)

}
