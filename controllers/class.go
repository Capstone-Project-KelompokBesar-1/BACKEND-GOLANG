package controllers

import (
	"net/http"
	"ourgym/dto"
	"ourgym/helpers"
	"ourgym/services"

	"github.com/labstack/echo/v4"
)

type ClassController struct {
	classService services.ClassService
}

func NewClassController(classService services.ClassService) *ClassController {
	return &ClassController{
		classService,
	}
}

func (uc *ClassController) GetAll(c echo.Context) error {
	name := c.QueryParam("name")

	classes := uc.classService.GetAll("", name)

	return c.JSON(http.StatusOK, Response(http.StatusOK, "Success get classes", classes))
}

func (uc *ClassController) GetAllOnlineClass(c echo.Context) error {
	name := c.QueryParam("name")

	classes := uc.classService.GetAll("online", name)

	return c.JSON(http.StatusOK, Response(http.StatusOK, "Success get online classes", classes))
}

func (uc *ClassController) GetAllOfflineClass(c echo.Context) error {
	name := c.QueryParam("name")

	classes := uc.classService.GetAll("offline", name)

	return c.JSON(http.StatusOK, Response(http.StatusOK, "Success get offline classes", classes))
}

func (uc *ClassController) GetByID(c echo.Context) error {
	var id string = c.Param("id")

	class := uc.classService.GetByID(id)

	if class.ID == 0 {
		return c.JSON(http.StatusNotFound, Response(http.StatusNotFound, "Class Not Found", nil))
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "Class Found", class))
}

func (uc *ClassController) Create(c echo.Context) error {
	input := dto.ClassRequest{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, "Request invalid", nil))
	}

	if err := input.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, "Request invalid", nil))
	}

	thumbnail, err := c.FormFile("thumbnail")

	if err == nil {
		url := helpers.UploadImage(thumbnail, "thumbnail")
		if url != "" {
			input.Thumbnail = url
		}
	}

	class := uc.classService.Create(input)

	return c.JSON(http.StatusOK, Response(http.StatusOK, "Success create class", class))
}

func (uc *ClassController) Update(c echo.Context) error {
	input := dto.ClassRequest{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, "Request invalid", nil))
	}

	if err := input.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, Response(http.StatusBadRequest, "Request invalid", nil))
	}

	thumbnail, err := c.FormFile("thumbnail")

	if err == nil {
		url := helpers.UploadImage(thumbnail, "thumbnail")
		if url != "" {
			input.Thumbnail = url
		}
	}

	var classId string = c.Param("id")

	class := uc.classService.Update(classId, input)

	return c.JSON(http.StatusOK, Response(http.StatusOK, "Success Update Class", class))
}

func (uc *ClassController) Delete(c echo.Context) error {
	var classId string = c.Param("id")

	isSuccess := uc.classService.Delete(classId)

	if !isSuccess {
		return c.JSON(http.StatusNotFound, Response(http.StatusNotFound, "Class Not Found", nil))
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "Class Success Deleted", nil))
}

func (uc *ClassController) DeleteMany(c echo.Context) error {
	ids := c.QueryParam("ids")

	isSuccess := uc.classService.DeleteMany(ids)

	if !isSuccess {
		return c.JSON(http.StatusNotFound, Response(http.StatusNotFound, "Classes Not Found", nil))
	}

	return c.JSON(http.StatusOK, Response(http.StatusOK, "Classes Success Deleted", nil))
}
