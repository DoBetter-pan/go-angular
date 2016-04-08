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
    "strings"
	"io/ioutil"
	"html/template"
)

var s_data map[string]string = map[string]string {"1":`
{
	"id": "1",
	"title": "Cookie1",
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
	"title": "Cookie2",
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

var s_retCode map[string] string = map[string] string{"success":`{
    "status": 1,
    "message": "successful"
}`, "error":`{
    "status": 0,
    "message": "unsuccessful"
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
	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	id := ""
	if len(parts) > 1 {
		id = parts[1]
	} else {
        id = "1"
    }
	fmt.Fprint(w, s_data[id])
}

func (controller *RecipeController) New(w http.ResponseWriter, r *http.Request) {
     tmpl, err := template.ParseFiles("app/content/index.html")
    if err != nil {
        log.Fatal("MainController::RenderMainFrame: ", err)
    }

    err = tmpl.Execute(w, controller)   
}

func (controller *RecipeController) Update(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")
	id := ""
	if len(parts) > 1 {
		id = parts[1]
	} else {
        id = "1"
    }

    //r.ParseForm()
    defer r.Body.Close()
    data, err := ioutil.ReadAll(r.Body)
    if err == nil {
        s_data[id] = string(data)
    }
	fmt.Fprint(w, s_data[id])
}

func (controller *RecipeController) Delete(w http.ResponseWriter, r *http.Request) {
     tmpl, err := template.ParseFiles("app/content/index.html")
    if err != nil {
        log.Fatal("MainController::RenderMainFrame: ", err)
    }

    err = tmpl.Execute(w, controller)   
}

