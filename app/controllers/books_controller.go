package controllers

import (
	"github.com/spf13/cast"
	"mumu/app/services"
	"mumu/pkg/helpers"
	"mumu/pkg/request"
	"mumu/pkg/route"
	"mumu/pkg/view"
	"net/http"
)

type BooksController struct {
	BaseController
}

// Index 文章列表页
func (bk *BooksController) Index(w http.ResponseWriter, r *http.Request) {
	channel := request.Query("g", "1", r)

	books := services.NewBook().GetHomeBooks(50, cast.ToInt(channel))
	view.Render(w, view.D{
		"g":     channel,
		"books": books,
	}, "index")
}

func (bk *BooksController) All(w http.ResponseWriter, r *http.Request) {
	channel := request.Query("g", "0", r)
	category := request.Query("c", "0", r)
	status := request.Query("s", "0", r)
	categories := services.NewBook().FindCat(cast.ToInt(channel))
	view.Render(w, view.D{
		"categories": categories,
		"channel":    channel,
		"c":          cast.ToInt(category),
		"status":     cast.ToInt(status),
	}, "all")
}

func (bk *BooksController) AjaxGet(w http.ResponseWriter, r *http.Request) {
	channel := request.Query("g", "0", r)
	cid := request.Query("c", "0", r)
	status := request.Query("status", "0", r)
	page := request.Query("page", "0", r)
	ret := services.NewBook().FindByAttr(cast.ToInt(cid), cast.ToInt(channel), cast.ToInt(status), cast.ToInt(page))

	helpers.ResponseJson(w, ret)
}

func (bk *BooksController) Detail(w http.ResponseWriter, r *http.Request) {
	id := cast.ToInt(route.GetRouteVariable("id", r))
	info := services.NewBook().Detail(uint64(id))
	view.Render(w, info, "detail")
}

func (bk *BooksController) DisplayChapter(w http.ResponseWriter, r *http.Request) {
	bookId := cast.ToUint64(route.GetRouteVariable("bookId", r))
	rolls, book := services.NewBook().ChapterRoll(bookId)
	view.Render(w, view.D{
		"BookId": bookId,
		"Rolls":  rolls,
		"Book":   book,
	}, "chapter")
}

func (bk *BooksController) AjaxChapter(w http.ResponseWriter, r *http.Request) {
	params := services.PageChapterParams{
		BookId:  cast.ToInt(request.Query("bookId", "0", r)),
		Page:    cast.ToInt(request.Query("page", "0", r)),
		OrderBy: request.Query("orderby", "0", r),
	}
	ret := services.NewBook().PageChapter(params)
	helpers.ResponseJson(w, ret)
}

func (bk *BooksController) Read(w http.ResponseWriter, r *http.Request) {
	chapterId := cast.ToInt(route.GetRouteVariable("chapterId", r))
	info := services.NewBook().Read(chapterId)
	view.Render(w, info, "read")
}

func (bk *BooksController) DisplaySearch(w http.ResponseWriter, r *http.Request) {
	data := services.NewBook().DisplaySearch()
	view.Render(w, data, "search")
}

func (bk *BooksController) Search(w http.ResponseWriter, r *http.Request) {
	page := cast.ToInt64(request.Query("page", "1", r))
	key := request.Query("key", "", r)
	ret := services.NewBook().Search(key, page)
	helpers.ResponseJson(w, ret)
}
