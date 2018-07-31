package admin

import (
	"fmt"
	"github.com/go-clog/clog"
	"github.com/julienschmidt/httprouter"
	"github.com/wuleying/silver-framework/admin/handlers"
	"github.com/wuleying/silver-framework/config"
	"github.com/wuleying/silver-framework/exceptions"
	"github.com/wuleying/silver-framework/utils"
	"net/http"
)

// HTTP struct
type HTTP struct {
	Config *config.Config
}

// Init 初始化Server
func (h *HTTP) Init() {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Oh no, not found"))
	})
	router.GET("/", handlers.Home)
	router.GET("/user", handlers.User)

	clog.Info(
		"Hello, %s. %s:%s, Goid: %d",
		h.Config.Setting["project_name"],
		h.Config.Setting["host"],
		h.Config.Setting["port"],
		utils.GetGoID())

	err := http.ListenAndServe(fmt.Sprintf(":%s", h.Config.Setting["port"]), router)
	exceptions.CheckError(err)

}
