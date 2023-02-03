package router

import (
	"richard-here/soluix/product-api/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/product")

	v1.Get("/", handler.ProductHandler.GetProductsHandler)
	v1.Post("/", handler.ProductHandler.AddProductHandler)
	v1.Put("/:code", handler.ProductHandler.EditProductHandler)
}
