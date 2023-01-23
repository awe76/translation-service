package mappers

import (
	"github.com/HqOapp/translation-service/src/inputs"
	"github.com/HqOapp/translation-service/src/models"
)

func InputToModel(input inputs.TranslationSet) []models.Translation {
	var (
		res = make([]models.Translation, 0)
	)

	for uuid, t := range input {
		for locale, content := range t {
			res = append(res, models.Translation{
				UUID:    uuid,
				Locale:  locale,
				Content: content,
			})
		}
	}

	return res
}
