package controller

import (
	"net/http"

	"github.com/b6217167/sa-64-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /right_treatment
func CreateRight_Treatment(c *gin.Context) {
	var right_treatment entity.Right_Treatment
	if err := c.ShouldBindJSON(&right_treatment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&right_treatment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": right_treatment})
}

// GET /right_treatment/:id
func GetRight_Treatment(c *gin.Context) {
	var right_treatment entity.Right_Treatment
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM right_treatments WHERE id = ?", id).Find(&right_treatment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": right_treatment})
}

// GET /right_treatments
func ListRight_Treatments(c *gin.Context) {
	var right_treatments []entity.Right_Treatment
	if err := entity.DB().Raw("SELECT * FROM right_treatments").Find(&right_treatments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": right_treatments})
}
