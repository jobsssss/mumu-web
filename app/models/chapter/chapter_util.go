package chapter

import (
	"github.com/repeale/fp-go"
	"github.com/spf13/cast"
	"mumu/pkg/model"
)

func Last(bookId uint64) (ret BookChapter) {
	condition := make(map[string]interface{})
	condition["book_id"] = bookId

	model.DB.Where(condition).Order("display_order DESC").First(&ret)
	return ret
}

func First(bookId uint64) (ret BookChapter) {
	condition := make(map[string]interface{})
	condition["book_id"] = bookId

	model.DB.Where(condition).Order("display_order ASC").First(&ret)
	return ret
}

func FirstByAttr(attrs map[string]interface{}) (wanted BookChapter) {
	model.DB.Where(attrs).First(&wanted)
	return
}

func FirstPre(c BookChapter) (wanted BookChapter) {
	model.DB.Where("display_order < ? and deleted_at = ? and book_id = ?", c.DisplayOrder, 0, c.BookId).
		Order("display_order DESC").First(&wanted)
	return
}

func FirstNext(c BookChapter) (wanted BookChapter) {
	model.DB.Where("display_order > ? and deleted_at = ? and book_id = ?", c.DisplayOrder, 0, c.BookId).
		Order("display_order ASC").First(&wanted)
	return
}

func PageFind(bookId int, page int, orderBy string) (ret []map[string]interface{}) {
	if orderBy == "" {
		orderBy = "ASC"
	}
	condition := make(map[string]interface{})
	condition["book_id"] = bookId
	offset := (page - 1) * 100

	var chapters []BookChapter
	model.DB.Where(condition).Order("display_order " + orderBy).Offset(offset).Limit(100).
		Find(&chapters)

	ret = fp.Map(func(t BookChapter) map[string]interface{} {
		return map[string]interface{}{
			"ID":     t.ID,
			"BookId": t.BookId,
			"Title":  t.Title,
			"Link":   "/read/" + cast.ToString(t.ID),
		}
	})(chapters)
	return
}

func Count(condition map[string]interface{}) (count int64) {
	model.DB.Model(&BookChapter{}).Where(condition).Count(&count)
	return
}
