package book

import (
	"mumu/app/models"
	"mumu/app/models/category"
)

const  (
	Loading = 1
	Finished = 2
)
type Book struct {
	models.BaseModel

	ChannelId   int8 `gorm:"column:channel_id;type:tinyint(1);not null;default:0"`
	Cid         uint
	Category    category.BookCategory `gorm:"foreignKey:Cid"`
	Name        string
	OldName     string
	Source      string
	SourceId    uint
	Author      string
	AuthorId    uint
	AuthorNote  string `gorm:"type:varchar(2048)"`
	Description string `gorm:"type:varchar(2048)"`
	Keywords    string `gorm:"type:varchar(1024)"`
	TotalPrice  uint
	WordsPrice          uint
	ChapterPrice        uint
	FreeChapters        uint
	Cover               string
	Status              uint
	OnlineStatus        uint
	Weight              uint
	IsVip               uint8
	IsBaoyue            uint8
	IsHot               uint8
	IsNew               uint8
	IsYy                uint8
	IsGreatest          uint8
	isGod               uint8
	IsSensitivity       uint8
	TotalWords          uint
	TotalChapters       uint
	TotalViews          uint
	TotalFavors         uint
	LastChapterId       uint
	LastChapterTime     uint
	ChapterUpdateRemind uint8
	AudioId             uint
	StoreType           uint8
	DeletedAt           uint
}
