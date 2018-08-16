package admin

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/go-clog/clog"
	"github.com/julienschmidt/httprouter"
	"github.com/wuleying/silver-framework/admin/handlers"
	"github.com/wuleying/silver-framework/exceptions"
	"github.com/wuleying/silver-framework/utils"
	"github.com/wuleying/silver-framework/uuid"
	"github.com/wuleying/silver-framework/version"
	"net/http"
)

// HTTP struct
type HTTP struct {
	Host  string
	Port  string
	UUID  snowflake.ID
	Error error
}

// Init 初始化Server
func (h *HTTP) Init() {
	// UUID
	h.UUID, h.Error = uuid.GetUUID()
	exceptions.CheckError(h.Error)

	router := httprouter.New()

	// not found
	router.NotFound = http.HandlerFunc(handlers.NotFound)

	// routes
	for _, route := range routes {
		router.Handle(route.Method, route.Pattern, route.Handle)
	}

	clog.Info(
		"%s:%s, version: %s, uuid: %s, goid: %d",
		h.Host,
		h.Port,
		version.Version,
		h.UUID.Base58(),
		utils.GetGoID())

	err := http.ListenAndServe(fmt.Sprintf("%s:%s", h.Host, h.Port), router)
	exceptions.CheckError(err)
}
