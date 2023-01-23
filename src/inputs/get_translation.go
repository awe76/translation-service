package inputs

type GetTranslationUri struct {
	UUID string `uri:"uuid" binding:"required,uuid"`
}

type GetTranslationQuery struct {
	Locale []string `form:"locale"`
}
