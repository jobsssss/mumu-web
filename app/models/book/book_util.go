package book

import (
	"github.com/meilisearch/meilisearch-go"
	"github.com/spf13/cast"
	"gorm.io/gorm/clause"
	"mumu/pkg/cache"
	"mumu/pkg/config"
	"mumu/pkg/helpers"
	"mumu/pkg/logger"
	"mumu/pkg/model"
	"time"
)

const (
	cacheDuration = time.Minute * 20
)

type BookInfo struct {
	ID              uint64
	Link            string
	Cover           string
	Title           string
	Desc            string
	Cat             string
	Status          string
	Views           uint
	ReadLink        string
	TotalWords      uint
	LastChapter     string
	LastChapterLink string
	UpdatedAt       string
}

func GetByChannelId(count int, gender int) (books []Book) {
	model.DB.Order("updated_at DESC").
		Where("channel_id = ?", gender).
		Preload("Category").
		Limit(count).
		Find(&books)
	return
}

func FindByAttr(cid int, status int, channelId int, count int, page int) (books []Book) {
	ql := model.DB.Order("id DESC")
	condition := make(map[string]interface{})
	if cid > 0 {
		condition["cid"] = cid
	}
	if status > 0 {
		condition["status"] = status
	}
	if channelId > 0 {
		condition["channel_id"] = channelId
	}
	offset := (page - 1) * count
	ql.Where(condition).
		Preload("Category").
		Offset(offset).
		Limit(count).
		Find(&books)
	return
}

func FirstById(id uint64) (ret Book) {
	model.DB.Where("id = ?", id).First(&ret)
	return
}

func FindByRand(count int) (data []BookInfo) {
	var ret []Book
	model.DB.Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "rand()"},
	}).Limit(count).Find(&ret)

	data = make([]BookInfo, len(ret))
	for k, v := range ret {
		data[k] = BookInfo{
			ID:    v.ID,
			Link:  "/detail/" + cast.ToString(v.ID),
			Cover: helpers.OssUrl(v.Cover),
			Title: v.Name,
			Desc:  v.Description,
		}
	}
	return
}

func FirstByAttrs(attrs map[string]interface{}) (wanted Book) {
	model.DB.Where(attrs).First(&wanted)
	return
}

var searchClient = meilisearch.NewClient(meilisearch.ClientConfig{
	Host:   cast.ToString(config.Env("MEILISEARCH_HOST", "")),
	APIKey: cast.ToString(config.Env("MEILISEARCH_KEY", "")),
})

func Search(key string, offset int64, limit int64) []interface{} {
	resp, err := searchClient.Index("book").Search(key, &meilisearch.SearchRequest{
		Offset: offset,
		Limit:  limit,
	})
	logger.LogIf(err)

	return resp.Hits
}

func All() []Book {
	var infos []Book
	model.DB.Order("id desc").Find(&infos)
	return infos
}

func SiteMap() (books []Book) {
	cacheKey := "book:sitemap"
	cache.GetObject(cacheKey, &books)
	if len(books) < 1 {
		books = All()
		cache.Set(cacheKey, books, cacheDuration)
	}
	return
}
