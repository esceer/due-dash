package api

import (
	"net/http"
	"strconv"

	"github.com/esceer/due-dash/backend/internal/api/model"
	"github.com/esceer/due-dash/backend/internal/service"
	"github.com/labstack/echo/v4"
)

type TaskApi interface {
	GetAll(c echo.Context) error
	GetById(c echo.Context) error
	Create(c echo.Context) error
	GenerateFromTemplate(c echo.Context) error
	UpdateStatus(c echo.Context) error
	Delete(c echo.Context) error
}

type taskApi struct {
	service service.TaskService
}

func NewTaskApiHandler(s service.TaskService) TaskApi {
	return &taskApi{
		service: s,
	}
}

func (a *taskApi) GetAll(c echo.Context) error {
	invoices, err := a.service.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, invoices)
}

func (a *taskApi) GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	invoice, err := a.service.GetById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, invoice)
}

func (a *taskApi) Create(c echo.Context) error {
	var invoice model.NewTask
	if err := c.Bind(&invoice); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := a.service.Create(&invoice); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}

func (a *taskApi) GenerateFromTemplate(c echo.Context) error {
	if err := a.service.GenerateFromTemplate(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}

func (a *taskApi) UpdateStatus(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var status model.Status
	if err := c.Bind(&status); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = a.service.UpdateStatus(id, string(status)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)

}

func (a *taskApi) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = a.service.Delete(id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
