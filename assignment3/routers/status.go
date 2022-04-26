package routers

import (
	"assignment3/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/update", controllers.UpdateJsonFile)
	router.GET("/status", controllers.GetStatusData)

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
