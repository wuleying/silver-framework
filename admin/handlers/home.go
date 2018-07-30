package handlers

import (
	"github.com/wuleying/silver-framework/exceptions"
	"github.com/wuleying/silver-framework/globals"
	"html/template"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

// Home 管理员首页
func Home(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	template, err := template.ParseFiles(globals.TemplateDir + "/home.html")
	exceptions.CheckError(err)

	template.Execute(response, nil)
}
