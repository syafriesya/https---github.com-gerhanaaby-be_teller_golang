package controllers

import (
	"golang/db"
	"golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostItem(c *gin.Context) {
	request := models.Item{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := db.GetDB().Debug().Create(&request).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Yeay Berhasil!",
	})
}
