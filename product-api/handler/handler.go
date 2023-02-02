package handler

import (
	"richard-here/soluix/product-api/database"
	"richard-here/soluix/product-api/repository"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Repo repository.Repo
}

var ProductHandler Handler

func InitHandler() {
	db := database.DB
	repo := repository.CreateRepository(db)
	ProductHandler = Handler{Repo: repo}
}

func (h *Handler) GetProductsHandler(c *fiber.Ctx) error {
	return nil
}
