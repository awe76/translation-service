package mappers

import (
	"github.com/HqOapp/translation-service/src/inputs"
	"github.com/HqOapp/translation-service/src/models"
)

func ModelToOutput(m []models.Translation) inputs.TranslationSet {
	var (
		res = inputs.TranslationSet{}
	)

	for _, t := range m {
		loc, ok := res[t.UUID]
		if !ok {
			loc = map[string]string{}
			res[t.UUID] = loc
		}

		loc[t.Locale] = t.Content
	}

	return res
}
