package admin

import (
	"fmt"
	"github.com/go-clog/clog"
	"github.com/wuleying/silver-framework/config"
	"github.com/wuleying/silver-framework/exceptions"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/wuleying/silver-framework/admin/handlers"
)

// HTTP struct
type HTTP struct {
	Config *config.Config
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

// Init 初始化Server
func (h *HTTP) Init() {
	router := httprouter.New()
	router.GET("/", handlers.Home)
	router.GET("/user", handlers.User)

	clog.Info(
		"Hello, %s. %s:%s",
		h.Config.Setting["project_name"],
		h.Config.Setting["host"],
		h.Config.Setting["port"],
	)

	err := http.ListenAndServe(fmt.Sprintf(":%s", h.Config.Setting["port"]), router)
	exceptions.CheckError(err)

}
