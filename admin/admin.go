package admin

import (
	"fmt"
	"github.com/go-clog/clog"
	"github.com/wuleying/silver-framework/admin/handlers"
	"github.com/wuleying/silver-framework/config"
	"github.com/wuleying/silver-framework/exceptions"
	"net/http"
)

// HTTP struct
type HTTP struct {
	Config *config.Config
}

// Init 初始化Server
func (h *HTTP) Init() {
	clog.Info("Hello, %s", h.Config.Setting["project_name"])

	http.HandleFunc("/", handlers.Home)

	err := http.ListenAndServe(fmt.Sprintf(":%s", "10100"), nil)
	exceptions.CheckError(err)
}
