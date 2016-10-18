/**
* @file loginSrvController.go
* @brief login service controller
* @author yingx
* @date 2015-02-27
 */

package controller

import (
	"net/http"
	_ "fmt"
    _ "strings"
    _ "strconv"
	"io/ioutil"
	model "go-angular/server/model"
)

type LoginSrvController struct {
}

func NewLoginSrvController() *LoginSrvController {
	return &LoginSrvController{}
}

func (controller *LoginSrvController) Query(w http.ResponseWriter, r *http.Request) {
    login := &model.LoginSrvModel{}
    res, err := login.FindAll()
    if err != nil {
        res = "[]"
    }

    SendBack(w, res)
}

func (controller *LoginSrvController) Get(w http.ResponseWriter, r *http.Request) {
    id := GetId(r)

    login := &model.LoginSrvModel{}
    res, err := login.Find(id)
    if err != nil {
        res = "{}"
    }

    SendBack(w, res)
}

func (controller *LoginSrvController) New(w http.ResponseWriter, r *http.Request) {
    res := "{}"

    //r.ParseForm()
    defer r.Body.Close()
    data, err := ioutil.ReadAll(r.Body)
    if err == nil {
        login := &model.LoginSrvModel{}
        res, err = login.Insert(string(data))
        if err != nil {
            res = "{}"
        }
    }

    SendBack(w, res)
}

func (controller *LoginSrvController) Update(w http.ResponseWriter, r *http.Request) {
    res := "{}"
    id := GetId(r)

    //r.ParseForm()
    defer r.Body.Close()
    data, err := ioutil.ReadAll(r.Body)
    if err == nil {
        login := &model.LoginSrvModel{}
        res, err = login.Update(id, string(data))
        if err != nil {
            res, err = login.Find(id)
            if err != nil {
                res = "{}"
            }
        }
    }

    SendBack(w, res)
}

func (controller *LoginSrvController) Delete(w http.ResponseWriter, r *http.Request) {
    id := GetId(r)

    login := &model.LoginSrvModel{}
    err := login.Delete(id)
    res := GetError(err)

    SendBack(w, res)
}

