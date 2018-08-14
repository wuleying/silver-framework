package admin

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/go-clog/clog"
	"github.com/julienschmidt/httprouter"
	"github.com/wuleying/silver-framework/admin/handlers"
	"github.com/wuleying/silver-framework/config"
	"github.com/wuleying/silver-framework/exceptions"
	"github.com/wuleying/silver-framework/utils"
	"github.com/wuleying/silver-framework/uuid"
	"github.com/wuleying/silver-framework/version"
	"net/http"
)

// HTTP struct
type HTTP struct {
	Config *config.Config
	UUID   snowflake.ID
}

// Init 初始化Server
func (h *HTTP) Init() {
	// UUID
	h.UUID = uuid.GetUUID()

	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Oh no, not found"))
	})

	router.GET("/", handlers.Home)
	router.GET("/user", handlers.User)
	router.GET("/metric", handlers.Metric)

	clog.Info(
		"Hello, %s. %s:%s, version: %s, uuid: %s, goid: %d",
		h.Config.Setting["project_name"],
		h.Config.Setting["host"],
		h.Config.Setting["port"],
		version.Version,
		h.UUID.Base58(),
		utils.GetGoID())

	err := http.ListenAndServe(fmt.Sprintf(":%s", h.Config.Setting["port"]), router)
	exceptions.CheckError(err)

}
