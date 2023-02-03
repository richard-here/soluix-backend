package handler

import (
	"log"
	"net/url"
	"richard-here/soluix/product-api/database"
	networkmodel "richard-here/soluix/product-api/model/network"
	"richard-here/soluix/product-api/repository"
	validate "richard-here/soluix/product-api/validator"
	"strconv"

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
	p := repository.Pagination{}

	lq, err := strconv.Atoi(c.Query("limit", "20"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	pq, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	nq := c.Query("name", "")
	scq := c.Query("subcategory", "")
	bq := c.Query("brand", "")
	minpq, err := strconv.Atoi(c.Query("minprice", "0"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	maxpq, err := strconv.Atoi(c.Query("maxprice", "0"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	boolq := c.Query("status", "")

	qs := new(networkmodel.GetQueries)
	qs.Limit = lq
	qs.Page = pq
	qs.Sort = url.QueryEscape(c.Query("sort", "id desc"))
	qs.Name = nq
	qs.Subcategory = scq
	qs.Brand = bq
	qs.MinPrice = minpq
	qs.MaxPrice = maxpq
	qs.Status = boolq

	errs := validate.ValidateGetQueries(*qs)
	log.Println(errs)
	if errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": errs,
		})
	}

	p.Limit = qs.Limit
	p.Page = qs.Page
	p.Sort, _ = url.QueryUnescape(qs.Sort)

	pagination, err := h.Repo.GetProducts(p, qs.Name, qs.Subcategory, qs.Brand, qs.MinPrice, qs.MaxPrice, qs.Status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   pagination,
	})
}

func (h *Handler) AddProductHandler(c *fiber.Ctx) error {
	ab := new(networkmodel.AddBody)

	err := c.BodyParser(ab)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	errs := validate.ValidateAddBody(*ab)
	if errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"errors": errs,
		})
	}

	lp := ab.ToLocalModel()
	err = h.Repo.AddProduct(lp)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Product was added",
		"data":    lp,
	})
}

func (h *Handler) EditProductHandler(c *fiber.Ctx) error {
	ub := new(networkmodel.UpdateBody)

	err := c.BodyParser(ub)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	errs := validate.ValidateUpdateBody(*ub)
	if errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"errors": errs,
		})
	}

	code := c.Params("code")
	p, err := h.Repo.GetProduct(code)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	err = h.Repo.UpdateProduct(p, ub.Status, ub.RetailPrice)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Product was edited",
		"data":    p,
	})
}
