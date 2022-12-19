package controllers

import (
	"net/http"
	"ourgym/services"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	categoryService services.CategoryService
}

func NewCategoryController(categoryService services.CategoryService) *CategoryController {
	return &CategoryController{
		categoryService,
	}
}

func (cc *CategoryController) GetAll(c echo.Context) error {
	name := c.QueryParam("name")
	categories := cc.categoryService.GetAll(name)

	if len(categories) == 0 {
		return c.JSON(http.StatusNotFound, Response(http.StatusNotFound, "Categories Not Found", nil))
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "Success Get Categories", categories))
}

func (cc *CategoryController) GetByID(c echo.Context) error {
	var id string = c.Param("id")

	category := cc.categoryService.GetByID(id)

	if category.ID == 0 {
		return c.JSON(http.StatusNotFound, Response(http.StatusNotFound, "Category Not Found", nil))
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "Category Found", category))
}
