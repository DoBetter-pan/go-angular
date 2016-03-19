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
	"fmt"
	_ "io/ioutil"
	"html/template"
)

var s_data map[string]string = map[string]string {"1":`
{
	"id": "1",
	"title": "Cookies",
	"description": "Delicious, crisp on the outside, chewy on the outside, oozing with chocolatey goodness cookies. The best kind.",
	"ingredients": [
	{
		"amount": "1",
		"amountUnits": "packet",
		"ingredientName": "Chips Ahoy"
	}
	],
	"instructions": "1. Go buy a paket of Chips Ahoy\n2. Heat it up in an oven\n3. Enjoy warm cookies\n4. Learn how to bake cookies from somewhere else"
}`,"2":`
{
	"id": "2",
	"title": "Cookies",
	"description": "Delicious, crisp on the outside, chewy on the outside, oozing with chocolatey goodness cookies. The best kind.",
	"ingredients": [
	{
		"amount": "1",
		"amountUnits": "packet",
		"ingredientName": "Chips Ahoy"
	}
	],
	"instructions": "1. Go buy a paket of Chips Ahoy\n2. Heat it up in an oven\n3. Enjoy warm cookies\n4. Learn how to bake cookies from somewhere else"
}`}

type RecipeController struct {
}

func NewRecipeController() *RecipeController {
	return &RecipeController{}
}

func (controller *RecipeController) Query(w http.ResponseWriter, r *http.Request) {
	var recipes = "["
	for _, v := range s_data {
		recipes += v
		recipes += ","
	}
	data := recipes[:len(recipes) - 1] + "]"

	fmt.Fprint(w, data)
}

func (controller *RecipeController) Get(w http.ResponseWriter, r *http.Request) {
     tmpl, err := template.ParseFiles("app/content/index.html")
    if err != nil {
        log.Fatal("MainController::RenderMainFrame: ", err)
    }

    err = tmpl.Execute(w, controller)   
}

func (controller *RecipeController) New(w http.ResponseWriter, r *http.Request) {
     tmpl, err := template.ParseFiles("app/content/index.html")
    if err != nil {
        log.Fatal("MainController::RenderMainFrame: ", err)
    }

    err = tmpl.Execute(w, controller)   
}

func (controller *RecipeController) Update(w http.ResponseWriter, r *http.Request) {
     tmpl, err := template.ParseFiles("app/content/index.html")
    if err != nil {
        log.Fatal("MainController::RenderMainFrame: ", err)
    }

    err = tmpl.Execute(w, controller)   
}

func (controller *RecipeController) Delete(w http.ResponseWriter, r *http.Request) {
     tmpl, err := template.ParseFiles("app/content/index.html")
    if err != nil {
        log.Fatal("MainController::RenderMainFrame: ", err)
    }

    err = tmpl.Execute(w, controller)   
}

