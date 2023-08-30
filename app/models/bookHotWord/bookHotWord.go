package bookHotWord

import "mumu/app/models"

type BookHotWord struct {
	models.BaseModel

	Keyword      string
	DisplayOrder int
	DeletedAt    int
}
