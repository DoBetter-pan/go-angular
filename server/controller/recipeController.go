/**
* @file recipeController.go
* @brief recipe controller
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

type RecipeController struct {
}

func NewRecipeController() *RecipeController {
	return &RecipeController{}
}


func (controller *RecipeController) IndexAction(w http.ResponseWriter, r *http.Request) {
     tmpl, err := template.ParseFiles("client/app/recipe/index.html")
    if err != nil {
        log.Fatal("MainController::RenderMainFrame: ", err)
    }

    err = tmpl.Execute(w, controller) 
}
