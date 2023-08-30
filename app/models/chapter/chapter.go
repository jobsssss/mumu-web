package chapter

import (
	"mumu/app/models"
)

type BookChapter struct {
	models.BaseModel

	BookId       uint64
	Title        string
	Content      string
	Words        int
	AuthorNote   string
	IsVip        uint8
	DisplayOrder int
	Ip           string
	DeletedAt    int
}
