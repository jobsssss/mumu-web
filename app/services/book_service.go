package services

import (
	"math"
	"mumu/app/models/book"
	"mumu/app/models/bookHotWord"
	"mumu/app/models/category"
	"mumu/app/models/chapter"
	commentModel "mumu/app/models/comment"
	"mumu/pkg/cache"
	"mumu/pkg/helpers"
	"strings"
	"time"

	"github.com/repeale/fp-go"
	"github.com/spf13/cast"
)

type BookService struct {
	BaseService
}

const (
	cacheDuration = time.Minute * 20
)

func (bks *BookService) GetHomeBooks(count int, gender int) (books []book.Book) {
	cacheKey := cache.BooksHome + "_" + cast.ToString(gender)
	cache.GetObject(cacheKey, &books)

	if len(books) < 1 {
		books = book.GetByChannelId(count, gender)
		if len(books) > 0 {
			cache.Set(cacheKey, books, cacheDuration)
		}
	}
	return
}

func (bks *BookService) FindByAttr(cid int, channelId int, status int, page int) (ret map[string][]book.BookInfo) {
	// 获取分类
	books := book.FindByAttr(cid, status, channelId, 25, page)
	data := make([]book.BookInfo, len(books))
	for k, v := range books {
		statusStr := "完结"
		if v.Status == book.Loading {
			statusStr = "连载中"
		} else if v.Status == book.Finished {
			statusStr = "完结"
		}
		data[k] = book.BookInfo{
			ID:     v.ID,
			Link:   "detail/" + cast.ToString(v.ID),
			Cover:  helpers.OssUrl(v.Cover),
			Title:  v.Name,
			Desc:   v.Description,
			Cat:    v.Category.Title,
			Status: statusStr,
			Views:  v.TotalViews,
		}
	}
	ret = map[string][]book.BookInfo{"Lists": data}
	return
}

func (bks *BookService) FindCat(channelId int) (categories []category.BookCategory) {
	categories = category.All(channelId)
	return
}

func (bks *BookService) Detail(id uint64) (info map[string]interface{}) {
	bookInfo := book.FirstById(id)
	lchp := chapter.Last(bookInfo.ID)
	fchp := chapter.First(bookInfo.ID)
	comment, num := commentModel.FindByBookId(id, 10)
	likes := book.FindByRand(8)

	status := "休刊"
	if bookInfo.Status == book.Loading {
		status = "连载中"
	} else if bookInfo.Status == book.Finished {
		status = "已完结"
	}

	info = map[string]interface{}{
		"ID":            bookInfo.ID,
		"Cover":         helpers.OssUrl(bookInfo.Cover),
		"Desc":          bookInfo.Description,
		"Title":         bookInfo.Name,
		"KeyWord":       bookInfo.Keywords,
		"Author":        bookInfo.Author,
		"Views":         bookInfo.TotalViews,
		"LastChapter":   lchp.Title,
		"TotalWords":    bookInfo.TotalWords,
		"Status":        status,
		"LchpUpdatedAt": time.Unix(lchp.UpdatedAt, 0).Format("2006-01-02 03"),
		"LchpId":        lchp.ID,
		"LchpTitle":     lchp.Title,
		"FchpId":        fchp.ID,
		"Likes":         likes,
		"Comment":       comment,
		"CommentNum":    num,
	}

	return
}

type PageChapterParams struct {
	BookId  int
	Page    int
	OrderBy string
}

func (bks *BookService) PageChapter(p PageChapterParams) map[string]interface{} {
	ret := chapter.PageFind(p.BookId, p.Page, p.OrderBy)
	return map[string]interface{}{"lists": ret}
}

func (bks *BookService) ChapterRoll(bookId uint64) ([]map[string]interface{}, book.Book) {
	bookInfo := book.FirstById(bookId)
	countChapter := chapter.Count(map[string]interface{}{"book_id": bookInfo.ID, "deleted_at": 0})

	if countChapter == 0 {
		return nil, book.Book{}
	}

	chp := int(math.Ceil(float64(countChapter) / 100))
	rolls := make([]map[string]interface{}, chp)
	h := 0
	j := 1
	l := chp

	for i := 0; i < chp; i++ {
		k := (i + 1) + h*100
		start := 1
		if i > 0 {
			start = k - j
		}

		rolls[i] = map[string]interface{}{
			"Name": cast.ToString(start) + "~" + cast.ToString(k-j+100) + "章",
			"Asc":  i,
			"Desc": l - 1,
		}
		l--
		h++
		j++
	}
	return rolls, bookInfo
}

func (bks *BookService) Read(chapterId int) (ret map[string]interface{}) {
	// 获取book信息
	// 查询出上一个章和下一章的 chapterId
	chapterDetail := chapter.FirstByAttr(map[string]interface{}{"id": chapterId})
	preChapter := chapter.FirstPre(chapterDetail)
	nextChapter := chapter.FirstNext(chapterDetail)
	bookDetail := book.FirstByAttrs(map[string]interface{}{"id": chapterDetail.BookId})

	ret = map[string]interface{}{
		"ChapterId":     chapterId,
		"ChapterTitle":  chapterDetail.Title,
		"Content":       strings.Split(chapterDetail.Content, "\n"),
		"KeyWords":      chapterDetail.Content[0:200],
		"PreChapterId":  preChapter.ID,
		"NextChapterId": nextChapter.ID,
		"BookId":        bookDetail.ID,
		"BookTitle":     bookDetail.Name,
		"BookDesc":      bookDetail.Description,
	}

	return
}

func (bks *BookService) DisplaySearch() (data map[string]interface{}) {
	data = make(map[string]interface{})
	hots := bookHotWord.All()
	data["Hots"] = bookHotWord.All()
	var keywords string
	for _, v := range hots {
		keywords += v.Keyword + ","
	}
	data["Keywords"] = strings.TrimRight(keywords, ",")
	return
}

func (bks *BookService) Search(key string, page int64) map[string]interface{} {

	offset := (page - 1) * 10

	data := make(map[string]interface{})
	infos := book.Search(key, offset, 10)
	data["lists"] = fp.Map(func(t interface{}) map[string]interface{} {
		v := t.(map[string]interface{})
		return map[string]interface{}{
			"ID":          v["ID"],
			"Name":        v["Name"],
			"Description": v["Description"],
			"Status":      v["Status"],
			"TotalViews":  v["TotalViews"],
			"Cover":       helpers.OssUrl(cast.ToString(v["Cover"])),
		}
	})(infos)

	return data
}
