package admin

import (
	"fmt"
	"github.com/wuleying/silver-framework/admin/handlers"
	"github.com/wuleying/silver-framework/exceptions"
	"net/http"
)

// Server
func Server() {
	http.HandleFunc("/", handlers.Home)

	err := http.ListenAndServe(fmt.Sprintf(":%s", "10100"), nil)
	exceptions.CheckError(err)
}
