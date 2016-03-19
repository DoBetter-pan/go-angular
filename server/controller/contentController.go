/**
* @file contentController.go
* @brief content controller
* @author yingx
* @date 2015-02-27
 */

package controller

import (
	"net/http"
	"log"
	_ "fmt"
	_ "io/ioutil"
	"html/template"
)

type ContentController struct {
}

func NewContentController() *ContentController {
	return &ContentController{}
}


func (controller *ContentController) IndexAction(w http.ResponseWriter, r *http.Request) {
     tmpl, err := template.ParseFiles("client/app/content/index.html")
    if err != nil {
        log.Fatal("MainController::RenderMainFrame: ", err)
    }

    err = tmpl.Execute(w, controller)   
}
