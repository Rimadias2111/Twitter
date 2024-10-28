package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"project/models"
)

// @Security ApiKeyAuth
// @Router /v1/tweets/like/{tweet_id} [post]
// @Summary Like a tweet
// @Description API for liking a tweet
// @Tags tweet
// @Param tweet_id path string true "Tweet ID"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) LikeTweet(c *gin.Context) {
	tweetID := c.Param("tweet_id")

	parsedTweetID, err := uuid.Parse(tweetID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid UUID format of tweet: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	userIdStr, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			ErrorMessage: "User ID not found in context",
			ErrorCode:    "Unauthorized",
		})
		return
	}

	userID, err := uuid.Parse(userIdStr.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid UUID format from token: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	like := models.Like{
		ID:      uuid.New(),
		UserID:  userID,
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
// @Router /v1/tweets/unlike/{tweet_id} [delete]
// @Summary Unlike a tweet
// @Description API for unliking a tweet
// @Tags tweet
// @Param tweet_id path string true "Tweet ID"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) UnlikeTweet(c *gin.Context) {
	tweetID := c.Param("tweet_id")

	parsedTweetID, err := uuid.Parse(tweetID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid UUID format of tweet: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	userIdStr, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			ErrorMessage: "User ID not found in context",
			ErrorCode:    "Unauthorized",
		})
		return
	}

	userID, err := uuid.Parse(userIdStr.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid UUID format from token: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	if err := h.store.Like().Delete(userID, parsedTweetID); err != nil {
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
