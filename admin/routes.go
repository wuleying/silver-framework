package admin

import (
	"github.com/julienschmidt/httprouter"
	"github.com/wuleying/silver-framework/admin/handlers"
)

// Route
type Route struct {
	Method  string
	Pattern string
	Handle  httprouter.Handle
}

// Routes
type Routes []Route

// 路由规则
var routes = Routes{
	Route{"GET", "/", handlers.Home},
	Route{"GET", "/user", handlers.User},
	Route{"GET", "/metric", handlers.Metric},
}
