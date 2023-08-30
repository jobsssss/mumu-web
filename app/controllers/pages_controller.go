package controllers

import (
	"fmt"
	"github.com/spf13/cast"
	"mumu/app/models/book"
	"mumu/app/models/chapter"
	"mumu/pkg/config"
	"mumu/pkg/logger"
	"net/http"
	time2 "time"
)

// PagesController 处理静态页面
type PagesController struct {
}

// Home 首页
func (*PagesController) Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "<h1>Hello, 欢迎来到 沐沐悦读！</h1>")
	if err != nil {
		return
	}
}

// About 关于我们页面
func (*PagesController) About(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
		"<a href=\"mailto:summer@example.com\">test@example.com</a>")
	if err != nil {
		return
	}
}

// NotFound 404 页面
func (*PagesController) NotFound(w http.ResponseWriter, r *http.Request) {
	//w.WriteHeader(http.StatusNotFound)
	//_, err := fmt.Fprint(w, "<h1>请求页面未找到 :(</h1><p>如有疑惑，请联系我们。</p>")
	//if err != nil {
	//	return
	//}
	http.Redirect(w, r, "/", 301)
}

func (*PagesController) Telescope(w http.ResponseWriter, r *http.Request) {
	chapter.PageFind(352, 1, "ASC")
}

func (bk *PagesController) SiteMapForBook(w http.ResponseWriter, r *http.Request) {
	books := book.SiteMap()
	xmlfirst := "<?xml version=\"1.0\" encoding=\"utf-8\"?>\n    <urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\" xmlns:mobile=\"http://www.baidu.com/schemas/sitemap-mobile/1/\">"
	var xmlBody string
	for _, b := range books {
		url := config.GetString("app.url") + "/detail/" + cast.ToString(b.ID)
		time := time2.Unix(b.UpdatedAt, 0).Format(time2.RFC3339)
		xmlBody += "" +
			"<url>" +
			"<loc>" + url + "</loc>" +
			"<lastmod>" + time + "</lastmod>" +
			"</url>"
	}
	xmlend := "</urlset>"

	xml := xmlfirst + xmlBody + xmlend
	w.Header().Set("Content-Type", "text/xml")
	_, err := fmt.Fprintf(w, xml)
	logger.LogIf(err)
}
