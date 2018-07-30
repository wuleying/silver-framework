package handlers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/wuleying/silver-framework/globals"
	"github.com/wuleying/silver-framework/exceptions"
	"net/http"
	"html/template"
)

// User 用户首页
func User(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	template, err := template.ParseFiles(globals.TemplateDir + "/user.html")
	exceptions.CheckError(err)

	template.Execute(response, nil)
}