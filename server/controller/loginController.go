/**
* @file loginController.go
* @brief login controller
* @author yingx
* @date 2015-02-27
 */

package controller

import (
    _ "fmt"
	"net/http"
	"log"
	"html/template"
	//model "go-angular/server/model"
)

type LoginMainParams struct {
    Stylesheets []string
    Javscripts []string
    Startup template.HTML
}

type LoginController struct {
}

func NewLoginController() *LoginController {
    controller := &LoginController{}
    return controller
}

func (controller *LoginController) IndexAction(w http.ResponseWriter, r *http.Request) {
    mainParams := &LoginMainParams{
        Stylesheets: []string {
            "../extensions/bootstrap-3.3.5/dist/css/bootstrap.min.css",
            "../app/login/styles/login.css" },
        Javscripts: []string {
            "../extensions/angular-1.5.0/angular.js",
            "../extensions/angular-1.5.0/angular-route.js",
            "../extensions/angular-1.5.0/angular-resource.js",
            "../app/login/scripts/directives/directives.js",
            "../app/login/scripts/services/services.js",
            "../app/login/scripts/login.js",
            "../app/login/scripts/controllers/login.js" },
        Startup : "" }

    tmpl, err := template.ParseFiles("client/app/login/login.html")
    if err != nil {
        log.Fatal("LoginController::IndexAction: ", err)
    }

    err = tmpl.Execute(w, mainParams)
}

