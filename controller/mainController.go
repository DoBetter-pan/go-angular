/**
* @file mainController.go
* @brief main controller
* @author yingx
* @date 2015-01-10
 */

package controller

import (
	"net/http"
	"log"
	_ "fmt"
	"bytes"
	"html/template"
)

type MainController struct {
	Title string
	Stylesheets []string
	Javscripts []string
	Banner template.HTML
	Content template.HTML
	Footer template.HTML
	Startup template.HTML
}

func LoadInfoFromTemplate(filename string) string {
	var b bytes.Buffer
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		log.Fatal("MainController::LoadInfoFromTemplate: ", err)
	}

	err = tmpl.Execute(&b, nil)
	return b.String()
}

func LoadHeaderFromTemplate() template.HTML {
	return template.HTML(LoadInfoFromTemplate("views/main/header.html"))
}

func LoadContentFromTemplate() template.HTML {
	return template.HTML(LoadInfoFromTemplate("views/main/content.html"))
}

func LoadFooterFromTemplate() template.HTML {
	return template.HTML(LoadInfoFromTemplate("views/main/footer.html"))
}

func NewMainController() *MainController {
	controller := &MainController{
		Title: "go-express",
		Stylesheets: []string {
			"../extensions/bootstrap-3.3.5/dist/css/bootstrap.min.css",
			"../assets/css/ie10-viewport-bug-workaround.css",
			"../css/normalize.css",
			"../css/main-style.css" },
		Javscripts: []string {
			"../js/jquery-1.11.3/jquery-1.11.3.min.js",
			"../extensions/bootstrap-3.3.5/dist/js/bootstrap.min.js" },
		Startup : "" }
	controller.Banner = LoadHeaderFromTemplate()
	controller.Content = LoadContentFromTemplate()
	controller.Footer = LoadFooterFromTemplate()
	return controller
}

func (controller *MainController) RenderMainFrame(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/main/main.html")
	if err != nil {
		log.Fatal("MainController::RenderMainFrame: ", err)
	}

	err = tmpl.Execute(w, controller)
}
