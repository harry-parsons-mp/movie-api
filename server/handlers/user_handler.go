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

type UserHandler struct {
	server *server.Server
}

func NewUserHandler(server *server.Server) *UserHandler {
	return &UserHandler{server: server}
}

func (h *UserHandler) List(c echo.Context) error {
	var user []models.User
	h.server.Repos.User.List(&user)

	//response
	res := responses.NewUsersResponse(user)
	return c.JSON(http.StatusOK, res)
}
func (h *UserHandler) Get(c echo.Context) error {
	var user models.User
	id := c.Param("id")

	h.server.Repos.User.Get(id, &user)
	if user.ID == 0 {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to retrieve user id: %v", id))
	}
	//response
	res := responses.NewUserResponse(&user)
	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) Create(c echo.Context) error {
	var req requests.UserRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if req.Username == "" {
		return c.JSON(http.StatusBadRequest, "username is required")
	}
	//create user
	user := models.User{
		Name:     req.Name,
		Username: req.Username,
	}

	err := h.server.Repos.User.Create(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to add user")
	}
	//response
	res := responses.NewUserResponse(&user)
	return c.JSON(http.StatusCreated, res)
}

func (h *UserHandler) Update(c echo.Context) error {
	ID := c.Param("id")

	var updateRequest requests.UserRequest
	if err := c.Bind(&updateRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// check user exists
	user := &models.User{}
	h.server.Repos.User.Get(ID, user)
	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, fmt.Sprintf("Failed to find user with id:%s", ID))
	}

	// Update the user
	user.Name = updateRequest.Name
	user.Username = updateRequest.Username
	err := h.server.Repos.User.Update(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("failed to update user, %v", err))
	}

	//return the updated user
	res := responses.NewUserResponse(user)
	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) Delete(c echo.Context) error {
	id := c.Param("id")

	var toDelete models.User
	h.server.Repos.User.Get(id, &toDelete)
	if toDelete.ID == 0 {
		c.JSON(http.StatusNotFound, fmt.Sprintf("failed to find user of id: %v to delete", id))
	}

	err := h.server.Repos.User.Delete(&toDelete)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to update user")
	}

	//response
	res := fmt.Sprintf("User of id: %v deleted sucessfully", id)
	return c.JSON(http.StatusOK, res)
}
