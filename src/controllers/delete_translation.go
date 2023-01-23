package controllers

import (
	"net/http"

	"github.com/HqOapp/translation-service/src/inputs"
	"github.com/HqOapp/translation-service/src/models"
	"github.com/gin-gonic/gin"
)

func DeleteTranslation(c *gin.Context) {
	var (
		uri inputs.GetTranslationUri
		tr  []models.Translation
		db  = models.DB
	)

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// TODO: Seems that batch delete doesn't work,
	// nevertheless it should be possible via raw queries.
	// Since number of locales are limited seems not a big problem.
	db.Where("uuid = $1", uri.UUID).Find(&tr)
	for _, t := range tr {
		db.Delete(&t)
	}

	c.JSON(http.StatusOK, gin.H{})
}
