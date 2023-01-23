package controllers

import (
	"net/http"

	"github.com/HqOapp/translation-service/src/inputs"
	"github.com/HqOapp/translation-service/src/mappers"
	"github.com/HqOapp/translation-service/src/models"
	"github.com/gin-gonic/gin"
)

func UpdateTranslation(c *gin.Context) {
	var (
		input inputs.TranslationSet
		db    = models.DB
	)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	translations := mappers.InputToModel(input)
	db.Exec("SELECT update_translations($1)", models.ToArray(translations, models.TranslationToString))

	c.JSON(http.StatusOK, gin.H{"data": input})
}
