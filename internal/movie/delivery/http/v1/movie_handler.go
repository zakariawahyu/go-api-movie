package v1

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-api-movie/internal/domain"
	"github.com/zakariawahyu/go-api-movie/internal/transport/request"
	"net/http"
	"strconv"
)

type movieHandler struct {
	movieUsecase domain.MovieUsecase
}

func NewMovieHandler(e *echo.Echo, movieUsecase domain.MovieUsecase) {
	handler := &movieHandler{
		movieUsecase: movieUsecase,
	}

	apiV1 := e.Group("/api/v1")
	apiV1.GET("/movie", handler.Fetch)
	apiV1.GET("/movie/:id", handler.GetByID)
	apiV1.POST("/movie", handler.Create)
	apiV1.PUT("/movie/:id", handler.Update)
	apiV1.DELETE("/movie/:id", handler.Delete)
}

func (h *movieHandler) Fetch(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.movieUsecase.Fetch(ctx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"data": res})
}

func (h *movieHandler) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	movie, err := h.movieUsecase.GetByID(ctx, int64(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"data": movie})
}

func (h *movieHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.CreateMovieRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return c.JSON(http.StatusBadRequest, errVal)
	}

	movie, err := h.movieUsecase.Create(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"data": movie})
}

func (h *movieHandler) Update(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))

	req := request.UpdateMovieRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return c.JSON(http.StatusBadRequest, errVal)
	}

	movie, err := h.movieUsecase.Update(ctx, req, int64(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"data": movie})
}

func (h *movieHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	if err := h.movieUsecase.Delete(ctx, int64(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "todo deleted",
	})
}
