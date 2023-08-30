package comment

import (
	"github.com/repeale/fp-go"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"mumu/pkg/helpers"
	"mumu/pkg/model"
)

func Find(condition map[string]interface{}, count int) (comment []BookComment, num int64) {
	model.DB.Where(condition).Order("id ASC").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("nickname", "uid", "avatar")
	}).Limit(count).Find(&comment)
	model.DB.Model(&BookComment{}).Where(condition).Count(&num)
	return
}

func FindByBookId(bookId uint64, count int) (lists []map[string]interface{}, num int64) {
	data, num := Find(map[string]interface{}{"book_id": bookId}, count)

	lists = fp.Map(func(v BookComment) map[string]interface{} {
		nickName := v.User.Nickname
		if nickName == "" {
			nickName = "书友-" + cast.ToString(v.ID)
		}
		return map[string]interface{}{
			"ID":       v.ID,
			"Content":  v.Content,
			"Avatar":   helpers.Asset("static/images/avatar.png"),
			"UserId":   v.User.Uid,
			"NickName": nickName,
		}
	})(data)

	return
}
