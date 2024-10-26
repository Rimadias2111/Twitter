package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/logger"
	"net/http"
	"project/database"
	"project/models"
	"strconv"
	"strings"
)

type Controller struct {
	store database.IStore
}

func NewController(store database.IStore) *Controller {
	return &Controller{store: store}
}

func ParsePageQueryParam(c *gin.Context) (uint64, error) {
	pageStr := c.Query("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.ParseUint(pageStr, 10, 30)
	if err != nil {
		return 0, err
	}
	if page == 0 {
		return 1, nil
	}
	return page, nil
}

func ParseLimitQueryParam(c *gin.Context) (uint64, error) {
	limitStr := c.Query("limit")
	if limitStr == "" {
		limitStr = "10"
	}
	limit, err := strconv.ParseUint(limitStr, 10, 30)
	if err != nil {
		return 0, err
	}

	if limit == 0 {
		return 10, nil
	}
	return limit, nil
}

func getUserInfo(h *Controller, c *gin.Context, accessibleUserTypes []string) error {
	var (
		ErrUnauthorized = errors.New("Unauthorized")
		accessToken     string
	)

	accessToken = c.GetHeader("Authorization")
	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			ErrorMessage: "Unauthorized request",
			ErrorCode:    "Unauthorized",
		})
		return errors.New("unauthorized request")
	}

	tokenArr := strings.Split(accessToken, " ")
	if len(tokenArr) == 2 {
		accessToken = tokenArr[1]
	}

	resp, err := h.service.Auth().CheckPermission(context.Background(), models.CheckPermissionRequest{
		AccessToken:         accessToken,
		AccessibleUserTypes: accessibleUserTypes,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			ErrorMessage: "Unauthorized request",
			ErrorCode:    config.ErrorCodeUnauthorized,
		})
		h.log.Error("Unauthorized request: ", logger.Error(err))
		return errors.New("unauthorized request")
	}

	if !resp.HasPermission {
		c.JSON(http.StatusForbidden, models.ResponseError{
			ErrorMessage: "Forbidden request",
			ErrorCode:    config.ErrorCodeForbidden,
		})
		h.log.Error("Request Forbidden")
		return errors.New("request Forbidden")
	}

	c.Set("userId", resp.UserId)
	c.Set("user_type_id", resp.UserTypeId)

	return nil
}
