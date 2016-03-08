/**
* @file angularController.go
* @brief angular controller
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

type AngularController struct {
}

func NewAngularController() *AngularController {
	return &AngularController{}
}

func LoadAngularIndexFromTemplate() template.HTML {
	return template.HTML(LoadInfoFromTemplate("views/angular/index.html"))
}

func (controller *AngularController) IndexAction(w http.ResponseWriter, r *http.Request) {
	startup := `
	<script type="text/javascript">
    $(function() {
        console.log("started....")
    });	
	</script>`

	mainConterller := NewMainController()
    /*
	//add new javascript and css
	mainConterller.Stylesheets = append(mainConterller.Stylesheets, []string{
        "../extensions/angular-1.5.0/togglemenu_boot/togglemenu.css",
        "../extensions/angular-1.5.0/angular-aside/css/angular-aside.min.css",
        "../extensions/angular-1.5.0/angular-busy/angular-busy.css",
        "../extensions/angular-1.5.0/toaster/toaster.css",
        "../css/MaterialIcons/material-icons.css",
        "../css/fontawesome/css/font-awesome.min.css"}...)
    mainConterller.Javscripts = append(mainConterller.Javscripts, []string{
        "../extensions/angular-1.5.0/angular.min.js",
        "../extensions/angular-1.5.0/ui-bootstrap-tpls-1.2.0.min.js",
        "../extensions/angular-1.5.0/angular-route.min.js",
        "../extensions/angular-1.5.0/togglemenu_boot/togglemenu.js",
        "../extensions/angular-1.5.0/angular-aside/js/angular-aside.min.js",
        "../extensions/angular-1.5.0/angular-sanitize.min.js",
        "../extensions/angular-1.5.0/angular-animate/angular-animate.min.js",
        "../extensions/angular-1.5.0/i18n/angular-locale_zh-cn.js",
        "../extensions/angular-1.5.0/angular-busy/angular-busy.js",
        "../extensions/angular-1.5.0/toaster/toaster.js"}...)
    */
	mainConterller.Startup = template.HTML(startup)
	mainConterller.Content = LoadAngularIndexFromTemplate()
	mainConterller.RenderMainFrame(w, r)
}
