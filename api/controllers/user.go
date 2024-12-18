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
// @Param user body models.CreateUser true "User data"
// @Success 200 {object} models.ResponseId
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) CreateUser(c *gin.Context) {
	var userModel models.CreateUser
	if err := c.ShouldBindJSON(&userModel); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while binding JSON: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	hashPassword, err := etc.GeneratePasswordHash(userModel.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while hashing password" + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}
	user := models.User{
		Name:         userModel.Name,
		Bio:          userModel.Bio,
		Username:     userModel.Username,
		Password:     string(hashPassword),
		ProfileImage: userModel.ProfileImage,
	}

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
// @Router /v1/users [put]
// @Summary Update a user
// @Description API for updating a user
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.UpdateUser true "User data"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) UpdateUser(c *gin.Context) {
	var userModel models.UpdateUser
	userIdStr, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			ErrorMessage: "User ID is not provided",
			ErrorCode:    "Unauthorized",
		})
		return
	}

	userId, err := uuid.Parse(userIdStr.(string))

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid user ID format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	if err := c.ShouldBindJSON(&userModel); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while binding JSON: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	user := models.User{
		Id:           userId,
		Name:         userModel.Name,
		Bio:          userModel.Bio,
		Username:     userModel.Username,
		ProfileImage: userModel.ProfileImage,
	}

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
// @Param user_id path string true "User ID"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) DeleteUser(c *gin.Context) {
	idStr := c.Param("user_id")
	id := uuid.MustParse(idStr)

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
	idStr := c.Param("user_id")
	id := uuid.MustParse(idStr)

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
// @Param id_followers query string false "id to get followers"
// @Param id_following query string false "id to get followings"
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
	followersQuery := c.Query("id_followers")
	followers, err := uuid.Parse(followersQuery) // Followers
	followingQuery := c.Query("id_followings")
	following, err := uuid.Parse(followingQuery) // Following

	req := models.GetAllUsersRequest{
		Page:      page,
		Limit:     limit,
		Search:    search,
		Followers: followers,
		Following: following,
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
