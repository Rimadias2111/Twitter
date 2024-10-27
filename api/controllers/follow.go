package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"project/models"
)

// @Security ApiKeyAuth
// @Router /v1/users/{user_id}/follow [post]
// @Summary Follow a user
// @Description API for following a user
// @Tags user
// @Param id path string true "User ID to follow"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) FollowUser(c *gin.Context) {
	followedID := c.Param("user_id")

	parsedFollowedID, err := uuid.Parse(followedID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid UUID format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	followerIdStr, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			ErrorMessage: "User ID not found in context",
			ErrorCode:    "Unauthorized",
		})
		return
	}

	followerID, err := uuid.Parse(followerIdStr.(string))

	isFollowing, err := h.store.Follow().IsFollowing(followerID, parsedFollowedID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while checking follow status: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	if isFollowing {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Already following this user",
			ErrorCode:    "Bad Request",
		})
		return
	}

	follow := models.Follow{
		ID:         uuid.New(),
		FollowerID: followerID,
		FollowedID: parsedFollowedID,
	}

	if err := h.store.Follow().Create(&follow); err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while following the user: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "User followed successfully",
	})
}

// @Security ApiKeyAuth
// @Router /v1/users/{user_id}/unfollow [delete]
// @Summary Unfollow a user
// @Description API for unfollowing a user
// @Tags user
// @Param id path string true "User ID to unfollow"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) UnfollowUser(c *gin.Context) {
	followedID := c.Param("user_id")

	parsedFollowedID, err := uuid.Parse(followedID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid UUID format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	followerIdStr, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			ErrorMessage: "User ID not found in context",
			ErrorCode:    "Unauthorized",
		})
		return
	}

	followerID, err := uuid.Parse(followerIdStr.(string))

	if err := h.store.Follow().Delete(followerID, parsedFollowedID); err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while unfollowing the user: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "User unfollowed successfully",
	})
}
