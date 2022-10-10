package controllers

import (
	"tecmentor-api/database"
	"tecmentor-api/models"
	"tecmentor-api/services"

	"github.com/gin-gonic/gin"
)

func Session(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Session

	err := c.ShouldBindJSON(&p)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON " + err.Error(),
		})
		return
	}

	var user models.Users
	dbError := db.Where("email = ?", p.Email).First(&user).Error

	if dbError != nil {
		c.JSON(400, gin.H{
			"error": "cannot find user",
		})
		return
	}

	if user.Password != services.Sha256Encoder(p.Password) {
		c.JSON(400, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	token, err := services.NewJWTService().GenerateToken(user.ID)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Create token error",
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})

}
