package api

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"project/api/controllers"
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
		//user endpoints
		api.POST("/users", cont.CreateUser)
		api.PUT("/users/:user_id", cont.UpdateUser)
		api.DELETE("/users/:user_id", cont.DeleteUser)
		api.GET("/users/:user_id", cont.GetUser)
		api.GET("/users", cont.GetAllUsers)

		//tweet endpoints
		api.POST("/tweets", cont.CreateTweet)
		api.PUT("/tweets/:tweet_id", cont.UpdateTweet)
		api.DELETE("/tweets/:tweet_id", cont.DeleteTweet)
		api.GET("/tweets/:tweet_id", cont.GetTweet)
		api.GET("/tweets", cont.GetAllTweets)
	}

	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
