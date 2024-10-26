package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"project/etc"
	"project/models"
)

// @Security ApiKeyAuth
// @Router /v1/users [post]
// @Summary Create a user
// @Description API for creating a new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 200 {object} models.ResponseId
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while binding JSON: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	hashPassword, err := etc.GeneratePasswordHash(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while hashing password" + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}
	user.Password = string(hashPassword)

	id, err := h.store.User().Create(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while creating a user: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseId{Id: id})
}

// @Security ApiKeyAuth
// @Router /v1/users/{user_id} [put]
// @Summary Update a user
// @Description API for updating a user
// @Tags user
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Param user body models.User true "User data"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) UpdateUser(c *gin.Context) {
	var user models.User
	id, err := uuid.Parse(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while binding JSON: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while binding JSON: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	user.Id = id
	if err := h.store.User().Update(&user); err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while updating the user: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "User updated successfully",
	})
}

// @Security ApiKeyAuth
// @Router /v1/users/{user_id} [delete]
// @Summary Delete a user
// @Description API for deleting a user
// @Tags user
// @Param id path string true "User ID"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) DeleteUser(c *gin.Context) {
	id := c.Param("user_id")

	err := h.store.User().Delete(models.RequestId{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while deleting the user: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "User deleted successfully",
	})
}

// @Security ApiKeyAuth
// @Router /v1/users/{user_id} [get]
// @Summary Get a user by ID
// @Description API for retrieving a user by ID
// @Tags user
// @Param user_id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) GetUser(c *gin.Context) {
	id := c.Param("user_id")

	user, err := h.store.User().Get(models.RequestId{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while retrieving the user: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Security ApiKeyAuth
// @Router /v1/users [get]
// @Summary Get all users
// @Description API for retrieving all users with pagination and search
// @Tags user
// @Param page query int false "Page number"
// @Param limit query int false "Number of users per page"
// @Param search query string false "Search term"
// @Success 200 {object} models.GetAllUsersResponse
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) GetAllUsers(c *gin.Context) {
	page, err := ParsePageQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid page: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid limit: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	search := c.Query("search")

	req := models.GetAllUsersRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	}

	users, err := h.store.User().GetAll(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while retrieving users: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}
