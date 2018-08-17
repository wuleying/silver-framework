package handlers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/wuleying/silver-framework/exceptions"
	"github.com/wuleying/silver-framework/globals"
	"html/template"
	"net/http"
)

// Metric 度量管理
func Metric(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	template, err := template.ParseFiles(globals.TemplateDir + "/metric.html")
	exceptions.CheckError(err)

	template.Execute(response, nil)
}
