package bootstrap

import (
	"embed"
	"github.com/gorilla/mux"
	"io/fs"
	"mumu/app/middlewares"
	"mumu/pkg/route"
	"mumu/routes"
	"net/http"
	//"net/http"
)

// SetupRoute 路由初始化
func SetupRoute(staticFS embed.FS) *mux.Router {
	router := mux.NewRouter()
	router.Use(middlewares.Recover)
	routes.RegisterWebRoutes(router)

	route.SetRoute(router)


	// 静态资源
	sub, _ := fs.Sub(staticFS, "public")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.FS(sub))))
	return router
}
