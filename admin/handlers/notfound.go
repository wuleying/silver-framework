package handlers

import (
	"github.com/wuleying/silver-framework/exceptions"
	"github.com/wuleying/silver-framework/globals"
	"html/template"
	"net/http"
)

func NotFound(response http.ResponseWriter, reuqest *http.Request) {
	template, err := template.ParseFiles(globals.TemplateDir + "/404.html")
	exceptions.CheckError(err)

	template.Execute(response, nil)
}
