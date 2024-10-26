package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"project/models"
)

// @Security ApiKeyAuth
// @Router /v1/tweets/{id}/like [post]
// @Summary Like a tweet
// @Description API for liking a tweet
// @Tags tweet
// @Param id path string true "Tweet ID"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) LikeTweet(c *gin.Context) {
	tweetID := c.Param("id")

	// Парсинг tweetID в UUID
	parsedTweetID, err := uuid.Parse(tweetID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid UUID format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	// Получение userID из токена
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			ErrorMessage: "User ID not found in context",
			ErrorCode:    "Unauthorized",
		})
		return
	}

	like := models.Like{
		ID:      uuid.New(),
		UserID:  userID.(uuid.UUID),
		TweetID: parsedTweetID,
	}

	if err := h.store.Like().Create(&like); err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while liking the tweet: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "Tweet liked successfully",
	})
}

// @Security ApiKeyAuth
// @Router /v1/tweets/{id}/unlike [delete]
// @Summary Unlike a tweet
// @Description API for unliking a tweet
// @Tags tweet
// @Param id path string true "Tweet ID"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) UnlikeTweet(c *gin.Context) {
	tweetID := c.Param("id")

	// Парсинг tweetID в UUID
	parsedTweetID, err := uuid.Parse(tweetID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid UUID format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	// Получение userID из токена
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			ErrorMessage: "User ID not found in context",
			ErrorCode:    "Unauthorized",
		})
		return
	}

	if err := h.store.Like().Delete(userID.(uuid.UUID), parsedTweetID); err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while unliking the tweet: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "Tweet unliked successfully",
	})
}
