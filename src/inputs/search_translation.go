package inputs

type SearchTranslationQuery struct {
	UUID          []string `form:"uuid"`
	Locale        []string `form:"locale"`
	SearchContent *string  `form:"content"`
}
