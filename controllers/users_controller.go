package controllers

import (
	"fmt"
	"tecmentor-api/database"
	"tecmentor-api/models"
	"tecmentor-api/services"

	"github.com/gin-gonic/gin"
)

func ShowUser(g *gin.Context) {
	g.JSON(200, gin.H{
		"teste": "primeiro endpoint",
	})
}

func CreateUser(g *gin.Context) {
	db := database.GetDatabase()

	var u models.Users

	err := g.ShouldBindJSON(&u)

	fmt.Println(u.Name)
	if err != nil {
		g.JSON(400, gin.H{
			"error": "cannot bind JSON " + err.Error(),
		})
		return
	}

	u.Password = services.Sha256Encoder(u.Password)
	fmt.Println(u.Password)

	err = db.Create(&u).Error
	if err != nil {
		g.JSON(400, gin.H{
			"error": "cannot create user " + err.Error(),
		})
		return
	}

	g.JSON(200, u)
}
