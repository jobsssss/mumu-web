package routes

import (
	"github.com/gorilla/mux"
	"mumu/app/controllers"
	"mumu/app/middlewares"
	"net/http"
)

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(r *mux.Router) {

	// 静态页面
	pc := new(controllers.PagesController)
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about")
	r.HandleFunc("/telescope", pc.Telescope).Methods("GET").Name("about")
	r.HandleFunc("/smpbook.xml", pc.SiteMapForBook).Methods("GET").Name("sitemap.book")

	// books
	bk := new(controllers.BooksController)
	r.HandleFunc("/", bk.Index).Methods("GET").Name("index")
	r.HandleFunc("/all", bk.All).Methods("GET").Name("all")
	r.HandleFunc("/book/ajax", bk.AjaxGet).Methods("GET").Name("book.ajax.get").Headers("X-Requested-With", "XMLHttpRequest") //.Headers("X-Requested-With", "XMLHttpRequest") 限定只能ajax访问
	r.HandleFunc("/detail/{id:[0-9]+}", bk.Detail).Methods("GET").Name("detail")
	r.HandleFunc("/chapter/{bookId}", bk.DisplayChapter).Methods("GET").Name("chapter")
	r.HandleFunc("/ajaxchapter", bk.AjaxChapter).Methods("GET").Name("ajaxchapter")
	r.HandleFunc("/read/{chapterId}", bk.Read).Methods("GET").Name("read")
	r.HandleFunc("/search", bk.DisplaySearch).Methods("GET").Name("search")
	r.HandleFunc("/ajaxsearch", bk.Search).Methods("GET").Name("ajaxsearch")

	// --- 全局中间件 ---
	// 开始会话
	r.Use(middlewares.StartSession)
}
