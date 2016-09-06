/**
* @file blogSrvController.go
* @brief blog service controller
* @author yingx
* @date 2015-02-27
 */

package controller

import (
	"net/http"
	"io/ioutil"
	model "go-angular/server/model"
)

type BlogSrvController struct {
}

func NewBlogSrvController() *BlogSrvController {
	return &BlogSrvController{}
}

func (controller *BlogSrvController) Query(w http.ResponseWriter, r *http.Request) {
    blog := &model.BlogSrvModel{}
    res, err := blog.FindAll()
    if err != nil {
        res = "[]"
    }

    SendBack(w, res)
}

func (controller *BlogSrvController) Get(w http.ResponseWriter, r *http.Request) {
    id := GetId(r)

    blog := &model.BlogSrvModel{}
    res, err := blog.Find(id)
    if err != nil {
        res = "{}"
    }

    SendBack(w, res)
}

func (controller *BlogSrvController) New(w http.ResponseWriter, r *http.Request) {
    res := "{}"

    //r.ParseForm()
    defer r.Body.Close()
    data, err := ioutil.ReadAll(r.Body)
    if err == nil {
        blog := &model.BlogSrvModel{}
        res, err = blog.Insert(string(data))
        if err != nil {
            res = "{}"
        }
    }

    SendBack(w, res)
}

func (controller *BlogSrvController) Update(w http.ResponseWriter, r *http.Request) {
    res := "{}"
    id := GetId(r)

    //r.ParseForm()
    defer r.Body.Close()
    data, err := ioutil.ReadAll(r.Body)
    if err == nil {
        blog := &model.BlogSrvModel{}
        res, err = blog.Update(id, string(data))
        if err != nil {
            res, err = blog.Find(id)
            if err != nil {
                res = "{}"
            }
        }
    }

    SendBack(w, res)
}

func (controller *BlogSrvController) Delete(w http.ResponseWriter, r *http.Request) {
    id := GetId(r)

    blog := &model.BlogSrvModel{}
    err := blog.Delete(id)
    res := GetError(err)

    SendBack(w, res)
}
