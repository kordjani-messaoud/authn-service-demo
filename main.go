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
	config.ExtractConfigParams(PATH, &config.GlobalConfigParams)

	app := fiber.New(fiber.Config{
		AppName:      "My Authn Service",
		ServerHeader: "Fiber",
	})
	middlewares.InitFiberMiddlewares(app, routes.InitPublicRoutes, nil)

	fmt.Println("Server listen on port:", config.GlobalConfigParams.ListenIP)

	err := app.Listen(fmt.Sprintf("%v:%v",
		config.GlobalConfigParams.ListenIP,
		config.GlobalConfigParams.ListenPort))

	log.Fatal(err)

}
