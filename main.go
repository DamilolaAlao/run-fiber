package main

import (
	// "fmt"
	"github.com/DamilolaAlao/run-fiber/controller"
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
)

func setupRoutes(app *fiber.App)  {
	// app.Get("/api/v1/", api.GetProducts)
	// api := app.Group("/api", cors())  // /api
	api := app.Group("/api")  // /api

//   v1 := api.Group("/v1", mysql())   // /api/v1
  v1 := api.Group("/v1")   // /api/v1
  v1.Get("/getproducts", controller.GetProducts)          // /api/v1/getproducts
//   v1.Get("/getMyProduct", api.GetMyProducts)          // /api/v1/getMyProduct

//   v2 := api.Group("/v2", mongodb()) // /api/v2
//   v2.Get("/list", handler)          // /api/v2/list
//   v2.Get("/user", handler) 
	// app.Get("/api/v1/book/:id", book.GetBook)
	// app.Post("/api/v1/book", book.NewBook)
	// app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func main()  {
	app := fiber.New()

	app.Use(logger.New())

	setupRoutes(app)

	app.Listen(8080)
}