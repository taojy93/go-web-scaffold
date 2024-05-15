package handlers

import (
	"net/http"
	"strconv"

	"go-web-scaffold/internal/errors"
	"go-web-scaffold/internal/models"
	"go-web-scaffold/internal/response"
	"go-web-scaffold/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.JSONResponse(c, http.StatusBadRequest, response.Error(errors.CodeInvalidParams, nil))
		return
	}

	createdUser, err := h.userService.CreateUser(&user)
	if err != nil {
		response.JSONResponse(c, http.StatusInternalServerError, response.Error(errors.CodeDatabaseError, nil))
		return
	}

	response.JSONResponse(c, http.StatusOK, response.Success(createdUser))
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.JSONResponse(c, http.StatusBadRequest, response.Error(errors.CodeInvalidParams, nil))
		return
	}

	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		response.JSONResponse(c, http.StatusInternalServerError, response.Error(errors.CodeDatabaseError, nil))
		return
	}

	response.JSONResponse(c, http.StatusOK, response.Success(user))
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		response.JSONResponse(c, http.StatusInternalServerError, response.Error(errors.CodeDatabaseError, nil))
		return
	}

	response.JSONResponse(c, http.StatusOK, response.Success(users))
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.JSONResponse(c, http.StatusBadRequest, response.Error(errors.CodeInvalidParams, nil))
		return
	}

	updatedUser, err := h.userService.UpdateUser(&user)
	if err != nil {
		response.JSONResponse(c, http.StatusInternalServerError, response.Error(errors.CodeDatabaseError, nil))
		return
	}

	response.JSONResponse(c, http.StatusOK, response.Success(updatedUser))
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.JSONResponse(c, http.StatusBadRequest, response.Error(errors.CodeInvalidParams, nil))
		return
	}

	err = h.userService.DeleteUser(uint(id))
	if err != nil {
		response.JSONResponse(c, http.StatusInternalServerError, response.Error(errors.CodeDatabaseError, nil))
		return
	}

	response.JSONResponse(c, http.StatusOK, response.Success(nil))
}
