package api

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"project/api/controllers"
	"project/api/middleware"
	_ "project/docs" //for swagger
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func Construct(cont controllers.Controller) *gin.Engine {
	r := gin.New()

	r.Static("/images", "./public/images")

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "pong"}) })

	api := r.Group("/v1")
	{
		//login
		api.POST("login", cont.LoginUser)

		//user endpoints
		api.POST("/users", cont.CreateUser)
		api.PUT("/users", middleware.AuthMiddleware(), cont.UpdateUser)
		api.DELETE("/users/:user_id", middleware.AuthMiddleware(), cont.DeleteUser)
		api.GET("/users/:user_id", cont.GetUser)
		api.GET("/users", cont.GetAllUsers)
		api.POST("/users/follow/:user_id", middleware.AuthMiddleware(), cont.FollowUser)
		api.DELETE("/users/unfollow/:user_id", middleware.AuthMiddleware(), cont.UnfollowUser)

		//tweet endpoints
		api.POST("/tweets", middleware.AuthMiddleware(), cont.CreateTweet)
		api.PUT("/tweets/:tweet_id", middleware.AuthMiddleware(), cont.UpdateTweet)
		api.DELETE("/tweets/:tweet_id", middleware.AuthMiddleware(), cont.DeleteTweet)
		api.GET("/tweets/:tweet_id", cont.GetTweet)
		api.GET("/tweets", cont.GetAllTweets)
		api.GET("/tweets/feed", middleware.AuthMiddleware(), cont.GetTweetsFeed)
		api.POST("/tweets/like/:tweet_id", middleware.AuthMiddleware(), cont.LikeTweet)
		api.DELETE("/tweets/unlike/:tweet_id", middleware.AuthMiddleware(), cont.UnlikeTweet)
		api.POST("/tweets/retweet/:tweet_id", middleware.AuthMiddleware(), cont.Retweet)
	}

	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
