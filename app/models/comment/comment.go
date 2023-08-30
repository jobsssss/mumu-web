package comment

import (
	"mumu/app/models"
	"mumu/app/models/book"
	"mumu/app/models/user"
)

type BookComment struct {
	models.BaseModel

	Uid         int
	User        user.User `gorm:"foreignKey:Uid"`
	BookId      uint64
	Book        book.Book `gorm:"foreignKey:BookId"`
	ChapterId   uint64
	Pid         uint64
	Content     string
	FromChannel uint8
	Status      uint8
	DeletedAt   int
}
