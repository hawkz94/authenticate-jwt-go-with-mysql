package routes

import (
	"tecmentor-api/controllers"
	"tecmentor-api/server/middlewares"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		users := main.Group("users")
		{
			users.POST("/create", controllers.CreateUser)
		}
		companies := main.Group("companies", middlewares.Auth())
		{
			companies.POST("/", controllers.CreateCompany)
		}

		main.POST("session", controllers.Session)
	}

	return router
}
