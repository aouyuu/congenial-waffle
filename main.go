package main

import (
	"github.com/aouyuu/thai-flood-radar/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/affacted", routes.GetAffactedAreas)
	r.GET("affected/overview", routes.GetAffactedAreasOverview)
	r.Run(":8080")
}
