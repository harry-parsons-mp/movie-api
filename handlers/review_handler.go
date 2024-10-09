package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"movie-api/models"
	"movie-api/repos"
	"movie-api/utils"
	"net/http"
)

type ReviewHandler struct {
	Repo repos.ReviewRepo
}

func NewReviewHandler(repo repos.ReviewRepo) *ReviewHandler {
	return &ReviewHandler{Repo: repo}
}

func (h *ReviewHandler) GetAllReviews(c echo.Context) error {
	reviews, err := h.Repo.GetAllReviews()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to retrieve reviews")
	}
	return c.JSON(http.StatusOK, reviews)
}
func (h *ReviewHandler) GetReviewByID(c echo.Context) error {
	id := utils.Str_to_uint(c.Param("id"))
	obj, err := h.Repo.GetReviewByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to retreive obj of id = %d", id))
	}
	return c.JSON(http.StatusOK, obj)

}
func (h *ReviewHandler) CreateReview(c echo.Context) error {
	var obj models.Review
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//title := c.FormValue("title")
	//if name != "" {
	//	obj.Title = title
	//}
	//content := c.FormValue("content")
	//if description != "" {
	//	obj.Description = description
	//}
	//genre := c.FormValue("genre")
	//if genre != "" {
	//	obj.Genre = genre
	//}

	err := h.Repo.CreateReview(&obj)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to add obj")
	}
	return c.JSON(http.StatusCreated, obj)
}

func (h *ReviewHandler) UpdateReview(c echo.Context) error {
	var obj models.Review
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err := h.Repo.UpdateReview(&obj)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to update obj")
	}
	return c.JSON(http.StatusOK, obj)
}

func (h *ReviewHandler) DeleteReview(c echo.Context) error {
	var obj models.Review
	if err := c.Bind(&obj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err := h.Repo.DeleteReview(&obj)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to update obj")
	}
	return c.JSON(http.StatusOK, obj)
}
