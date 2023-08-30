package category

import "mumu/app/models"

type BookCategory struct {
	models.BaseModel

	Title        string
	ChannelId    uint8
	Icon         string
	DisplayOrder uint
	DeletedAt    uint
}
