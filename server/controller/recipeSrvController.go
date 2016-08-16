/**
* @file recipeSrvController.go
* @brief recipe service controller
* @author yingx
* @date 2015-02-27
 */

package controller

import (
	"net/http"
	"fmt"
    "strings"
    "strconv"
	"io/ioutil"
	model "go-angular/server/model"
)

var s_retCode map[string] string = map[string] string{"success":`{
    "status": 1,
    "message": "successful"
}`, "error":`{
    "status": 0,
    "message": "unsuccessful"
}`}

type RecipeSrvController struct {
}

func NewRecipeSrvController() *RecipeSrvController {
	return &RecipeSrvController{}
}

func (controller *RecipeSrvController) Query(w http.ResponseWriter, r *http.Request) {
    recipe := &model.RecipeSrvModel{}
    res, err := recipe.FindAll()
    if err != nil {
        res = "[]"
    }

    w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	fmt.Fprint(w, res)
}

func (controller *RecipeSrvController) Get(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	var id int64 = 1
    if len(parts) > 1 {
        num, err := strconv.ParseInt(parts[1], 10, 64)
        if err != nil {
            num = 1
        }
        id = num
    }

    recipe := &model.RecipeSrvModel{}
    res, err := recipe.Find(id)
    if err != nil {
        res = "{}"
    }

    w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	fmt.Fprint(w, res)
}

func (controller *RecipeSrvController) New(w http.ResponseWriter, r *http.Request) {
    res := "{}"

    //r.ParseForm()
    defer r.Body.Close()
    data, err := ioutil.ReadAll(r.Body)
    if err == nil {
        recipe := &model.RecipeSrvModel{}
        res, err = recipe.Insert(string(data))
        if err != nil {
            res = "{}"
        }
    }

    w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	fmt.Fprint(w, res)
}

func (controller *RecipeSrvController) Update(w http.ResponseWriter, r *http.Request) {
    res := "{}"

	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	var id int64 = 1
    if len(parts) > 1 {
        num, err := strconv.ParseInt(parts[1], 10, 64)
        if err != nil {
            num = 1
        }
        id = num
    }

    //r.ParseForm()
    defer r.Body.Close()
    data, err := ioutil.ReadAll(r.Body)
    if err == nil {
        recipe := &model.RecipeSrvModel{}
        res, err = recipe.Update(id, string(data))
        if err != nil {
            res, err = recipe.Find(id)
            if err != nil {
                res = "{}"
            }
        }
    }

    w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
	fmt.Fprint(w, res)
}

func (controller *RecipeSrvController) Delete(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	var id int64 = 1
    if len(parts) > 1 {
        num, err := strconv.ParseInt(parts[1], 10, 64)
        if err != nil {
            num = 1
        }
        id = num
    }

    res := ""
    recipe := &model.RecipeSrvModel{}
    err := recipe.Delete(id)
    if err == nil {
        res = s_retCode["success"]
    } else {
        res = s_retCode["error"]
    }

	fmt.Fprint(w, res)
}

