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

	// not found
	router.NotFound = http.HandlerFunc(handlers.NotFound)

	// routes
	for _, route := range routes {
		router.Handle(route.Method, route.Pattern, route.Handle)
	}

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
