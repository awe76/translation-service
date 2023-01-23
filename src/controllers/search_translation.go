package controllers

import (
	"fmt"
	"net/http"

	"github.com/HqOapp/translation-service/src/inputs"
	"github.com/HqOapp/translation-service/src/mappers"
	"github.com/HqOapp/translation-service/src/models"
	"github.com/gin-gonic/gin"
)

func SearchTranslation(c *gin.Context) {
	var (
		q  inputs.SearchTranslationQuery
		tr []models.Translation
	)

	if err := c.ShouldBindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	clause := models.DB
	index := 1
	if len(q.UUID) > 0 {
		clause = clause.Where(fmt.Sprintf("uuid = ANY($%d)", index), q.UUID)
		index++
	}

	if len(q.Locale) > 0 {
		clause = clause.Where(fmt.Sprintf("locale = ANY($%d)", index), q.Locale)
		index++
	}

	if q.SearchContent != nil {
		clause = clause.Where(fmt.Sprintf("content LIKE $%d", index), fmt.Sprintf("%%%s%%", *q.SearchContent))
	}

	clause.Find(&tr)
	o := mappers.ModelToOutput(tr)
	c.JSON(http.StatusOK, gin.H{"data": o})

}
