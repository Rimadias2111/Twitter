package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"project/models"
)

// @Security ApiKeyAuth
// @Router /v1/tweets [post]
// @Summary Create a tweet
// @Description API for creating a new tweet
// @Tags tweet
// @Accept json
// @Produce json
// @Param tweet body models.Tweet true "Tweet data"
// @Success 200 {object} models.ResponseId
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) CreateTweet(c *gin.Context) {
	var tweet models.Tweet
	if err := c.ShouldBindJSON(&tweet); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while binding JSON: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	id, err := h.store.Tweet().Create(&tweet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while creating a tweet: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseId{Id: id})
}

// @Security ApiKeyAuth
// @Router /v1/tweets/{tweet_id} [put]
// @Summary Update a tweet
// @Description API for updating a tweet
// @Tags tweet
// @Accept json
// @Produce json
// @Param tweet_id path string true "Tweet ID"
// @Param tweet body models.Tweet true "Tweet data"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) UpdateTweet(c *gin.Context) {
	var tweet models.Tweet
	id := c.Param("tweet_id")

	parsedID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid UUID format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	if err := c.ShouldBindJSON(&tweet); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Error while binding JSON: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	tweet.Id = parsedID
	if err := h.store.Tweet().Update(&tweet); err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while updating the tweet: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "Tweet updated successfully",
	})
}

// @Security ApiKeyAuth
// @Router /v1/tweets/{tweet_id} [delete]
// @Summary Delete a tweet
// @Description API for deleting a tweet
// @Tags tweet
// @Param id path string true "Tweet ID"
// @Success 200 {object} models.ResponseSuccess
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) DeleteTweet(c *gin.Context) {
	id := c.Param("tweet_id")

	err := h.store.Tweet().Delete(models.RequestId{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while deleting the tweet: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseSuccess{
		Message: "Tweet deleted successfully",
	})
}

// @Security ApiKeyAuth
// @Router /v1/tweets/{id} [get]
// @Summary Get a tweet by ID
// @Description API for retrieving a tweet by ID
// @Tags tweet
// @Param tweet_id path string true "Tweet ID"
// @Success 200 {object} models.Tweet
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) GetTweet(c *gin.Context) {
	id := c.Param("tweet_id")

	_, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid UUID format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	tweet, err := h.store.Tweet().Get(models.RequestId{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while retrieving the tweet: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, tweet)
}

// @Security ApiKeyAuth
// @Router /v1/tweets [get]
// @Summary Get all tweets
// @Description API for retrieving all tweets with pagination and search
// @Tags tweet
// @Param page query int false "Page number"
// @Param limit query int false "Number of tweets per page"
// @Param search query string false "Search term"
// @Param user_id query string false "User ID for filtering tweets"
// @Success 200 {object} models.GetAllTweetsResponse
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) GetAllTweets(c *gin.Context) {
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

	userId := c.Query("user_id")

	search := c.Query("search")

	req := models.GetAllTweetsRequest{
		Limit:  limit,
		Page:   page,
		UserID: userId,
		Search: search,
	}

	tweets, err := h.store.Tweet().GetAll(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while retrieving tweets: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, tweets)
}

// @Security ApiKeyAuth
// @Router /v1/tweets/feed [get]
// @Summary Get tweets from followed users
// @Description API for retrieving tweets from users that the current user is following
// @Tags tweet
// @Param page query int false "Page number"
// @Param limit query int false "Number of tweets per page"
// @Success 200 {object} models.GetAllTweetsResponse
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) GetTweetsFeed(c *gin.Context) {
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

	req := models.GetAllTweetsRequest{
		Limit: limit,
		Page:  page,
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			ErrorMessage: "User ID not found in context",
			ErrorCode:    "Unauthorized",
		})
		return
	}

	tweets, err := h.store.Tweet().GetTweetsForUser(userID.(uuid.UUID), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while retrieving tweets feed: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, tweets)
}

// @Security ApiKeyAuth
// @Router /v1/tweets/{tweet_id}/retweet [post]
// @Summary Retweets a tweet
// @Description API for retweeting an existing tweet
// @Tags tweet
// @Param id path string true "Tweet ID to retweet"
// @Success 200 {object} models.ResponseId
// @Failure 400 {object} models.ResponseError "Invalid input"
// @Failure 500 {object} models.ResponseError "Internal server error"
func (h *Controller) Retweet(c *gin.Context) {
	id := c.Param("tweet_id")

	originalTweetID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			ErrorMessage: "Invalid UUID format: " + err.Error(),
			ErrorCode:    "Bad Request",
		})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			ErrorMessage: "User ID not found in context",
			ErrorCode:    "Unauthorized",
		})
		return
	}

	newTweet := models.Tweet{
		Id:        uuid.New(),
		UserID:    userID.(uuid.UUID),
		RetweetID: &originalTweetID,
	}

	retweetID, err := h.store.Tweet().Create(&newTweet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			ErrorMessage: "Error while creating a retweet: " + err.Error(),
			ErrorCode:    "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, models.ResponseId{Id: retweetID})
}