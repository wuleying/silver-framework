package admin

import (
	"fmt"
	"github.com/wuleying/silver-framework/admin/handlers"
	"github.com/wuleying/silver-framework/exceptions"
	"net/http"
)

type HTTP struct {
}

// Server
func (h *HTTP) Init() {
	http.HandleFunc("/", handlers.Home)

	err := http.ListenAndServe(fmt.Sprintf(":%s", "10100"), nil)
	exceptions.CheckError(err)
}
