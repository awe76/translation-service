package controllers

import (
	"net/http"

	"github.com/HqOapp/translation-service/src/inputs"
	"github.com/HqOapp/translation-service/src/mappers"
	"github.com/HqOapp/translation-service/src/models"
	"github.com/gin-gonic/gin"
)

// POST /translation
// Create a translation
func CreateTranslation(c *gin.Context) {
	var (
		input inputs.TranslationSet
	)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	translations := mappers.InputToModel(input)

	models.DB.Create(&translations)
	c.JSON(http.StatusCreated, gin.H{})
}
