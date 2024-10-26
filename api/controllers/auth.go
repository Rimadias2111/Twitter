package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/etc"
	auth "project/etc/jwt"
	"project/models"
)

// @Router /v1/login [post]
// @Summary User login
// @Description API for user login
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body models.LoginRequest true "User credentials"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 401 {object} models.ResponseError "Unauthorized"
// @Failure 500 {object} models.ResponseError "Internal Server Error"
func (h *Controller) LoginUser(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while binding JSON: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	user, err := h.store.User().GetByUsername(req.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			ErrorMessage: "User not found",
			ErrorCode:    "Unauthorized",
		})
		return
	}

	isPasswordValid, err := etc.CheckPassword(req.Password, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while checking password: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	if !isPasswordValid {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			ErrorMessage: "Invalid password",
			ErrorCode:    "Unauthorized",
		})
		return
	}

	token, err := auth.GenerateToken(user.Id.String(), user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while generating token: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.LoginResponse{
		Token: token,
	})
}
