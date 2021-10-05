package routes

import "github.com/gin-gonic/gin"

func GetAffactedAreas(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "affected areas"})
}

func GetAffactedAreasOverview(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "affected areas overview"})
}
