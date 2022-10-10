package controllers

import (
	"tecmentor-api/database"
	"tecmentor-api/models"

	"github.com/gin-gonic/gin"
)

func CreateCompany(g *gin.Context) {
	db := database.GetDatabase()

	var c models.Companies

	err := g.ShouldBindJSON(&c)

	if err != nil {
		g.JSON(400, gin.H{
			"error": "cannot bind JSON " + err.Error(),
		})
		return
	}

	err = db.Create(&c).Error
	if err != nil {
		g.JSON(400, gin.H{
			"error": "cannot create Company " + err.Error(),
		})
		return
	}

	g.JSON(200, c)
}
