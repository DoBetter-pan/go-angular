/**
* @file contentController.go
* @brief content controller
* @author yingx
* @date 2015-02-27
 */

package controller

import (
	"net/http"
	_ "log"
	_ "fmt"
	_ "io/ioutil"
	"html/template"
)

type ContentController struct {
}

func NewContentController() *ContentController {
	return &ContentController{}
}

func LoadContentIndexFromTemplate() template.HTML {
	return template.HTML(LoadInfoFromTemplate("server/views/content/index.html"))
}

func (controller *ContentController) IndexAction(w http.ResponseWriter, r *http.Request) {
	startup := `
	<script type="text/javascript">
    $(function() {
        console.log("started....")
    });	
	</script>`

	mainConterller := NewMainController()
	//add new javascript and css
	mainConterller.Stylesheets = append(mainConterller.Stylesheets, []string{
        "../content/styles/main.scss"}...)  
    mainConterller.Javscripts = append(mainConterller.Javscripts, []string{
        "../content/script/controllers/main.js",
        "../content/script/controllers/about.js",
        "../content/script/app.js"}...)
	mainConterller.Startup = template.HTML(startup)
	mainConterller.Content = LoadContentIndexFromTemplate()
	mainConterller.RenderMainFrame(w, r)
}
