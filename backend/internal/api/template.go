package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/esceer/due-dash/backend/internal/api/model"
	"github.com/esceer/due-dash/backend/internal/service"
	"github.com/labstack/echo/v4"
)

type TemplateApi interface {
	GetAll(c echo.Context) error
	GetById(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type templateApi struct {
	service service.TemplateService
}

func NewTemplateApiHandler(s service.TemplateService) TemplateApi {
	return &templateApi{
		service: s,
	}
}

func (a *templateApi) GetAll(c echo.Context) error {
	templates, err := a.service.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, templates)
}

func (a *templateApi) GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	template, err := a.service.GetById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, template)
}

func (a *templateApi) Create(c echo.Context) error {
	var template model.NewTemplate
	if err := c.Bind(&template); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := a.service.Create(&template); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}

func (a *templateApi) Update(c echo.Context) error {
	idFromPath, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var template model.Template
	if err := c.Bind(&template); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if idFromPath != template.Id {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("template.id mismatch"))
	}

	if err = a.service.Update(&template); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)

}

func (a *templateApi) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = a.service.Delete(id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
