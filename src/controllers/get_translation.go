package controllers

import (
	"net/http"

	"github.com/HqOapp/translation-service/src/inputs"
	"github.com/HqOapp/translation-service/src/mappers"
	"github.com/HqOapp/translation-service/src/models"
	"github.com/gin-gonic/gin"
)

func GetTranslation(c *gin.Context) {
	var (
		uri inputs.GetTranslationUri
		q   inputs.GetTranslationQuery
		tr  []models.Translation
	)

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	clause := models.DB.Where("uuid = $1", uri.UUID)

	if len(q.Locale) > 0 {
		clause = clause.Where("locale = ANY($2)", q.Locale)
	}

	clause.Find(&tr)

	o := mappers.ModelToOutput(tr)
	c.JSON(http.StatusOK, gin.H{"data": o})
}
