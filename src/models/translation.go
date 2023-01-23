package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Translation struct {
	gorm.Model
	UUID    string `gorm:"index:idx_uuid_locale,unique"`
	Locale  string `gorm:"index:idx_uuid_locale,unique"`
	Content string
}

func TranslationToString(t *Translation) string {
	return fmt.Sprintf("(0,,,,%s,%s,%s)", t.UUID, t.Locale, t.Content)
}
