package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"movie-api/models"
	"movie-api/server"
	"movie-api/server/requests"
	"movie-api/server/responses"
	"net/http"
)

type ReviewHandler struct {
	server *server.Server
}

func NewReviewHandler(server *server.Server) *ReviewHandler {
	return &ReviewHandler{server: server}
}

func (h *ReviewHandler) List(c echo.Context) error {
	var reviews []models.Review
	h.server.Repos.Review.List(&reviews)

	//response
	res := responses.NewReviewsResponse(reviews)
	return c.JSON(http.StatusOK, res)
}

func (h *ReviewHandler) GetByID(c echo.Context) error {
	id := c.Param("id")

	review := &models.Review{}
	h.server.Repos.Review.Get(id, review)
	if review.ID == 0 {
		return c.JSON(http.StatusNotFound, fmt.Sprintf("Failed to retreive review of id: %s", id))
	}

	//response
	res := responses.NewReviewResponse(review)
	return c.JSON(http.StatusOK, res)
}

func (h *ReviewHandler) Create(c echo.Context) error {
	var req requests.ReviewRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if req.Title == "" {
		c.JSON(http.StatusBadRequest, "review title is required")
	}

	// create review from request
	review := models.Review{
		Title:   req.Title,
		Content: req.Content,
		Score:   req.Score,
		UserID:  req.UserID,
		MovieID: req.MovieID,
	}
	err := h.server.Repos.Review.Create(&review)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to add review")
	}

	//response
	res := responses.NewReviewResponse(&review)
	return c.JSON(http.StatusCreated, res)
}

func (h *ReviewHandler) Update(c echo.Context) error {
	ID := c.Param("id")

	var req requests.ReviewRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// check review exists
	review := &models.Review{}
	h.server.Repos.Review.Get(ID, review)
	if review.ID == 0 {
		return c.JSON(http.StatusNotFound, fmt.Sprintf("Failed to find review with id:%s", ID))
	}

	// Update the review
	review.Title = req.Title
	review.Content = req.Content
	review.Score = req.Score
	review.UserID = req.UserID
	review.MovieID = req.MovieID

	err := h.server.Repos.Review.Update(review)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("failed to update review, %v", err))
	}

	// create response
	res := responses.NewReviewResponse(review)
	return c.JSON(http.StatusOK, res)
}

func (h *ReviewHandler) Delete(c echo.Context) error {
	var toDelete models.Review

	id := c.Param("id")
	h.server.Repos.Review.Get(id, &toDelete)
	if toDelete.ID == 0 {
		return c.JSON(http.StatusNotFound, fmt.Sprintf("failed to find review of id: %v to delete", id))
	}

	err := h.server.Repos.Review.Delete(&toDelete)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to update obj")
	}

	//response
	res := fmt.Sprintf("Review of id: %v deleted sucessfully", id)
	return c.JSON(http.StatusOK, res)
}
