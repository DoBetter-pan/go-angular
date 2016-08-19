/**
* @file basicSrvController.go
* @brief basic service controller
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
    "encoding/json"
	model "go-angular/server/model"
)

func GetId(r *http.Request) int64 {
	var id int64 = 1

	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")

    if len(parts) > 1 {
        num, err := strconv.ParseInt(parts[1], 10, 64)
        if err != nil {
            num = 1
        }
        id = num
    }

    return id
}

func SendBack(w http.ResponseWriter, data string) {
    //set header
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprint(w, data)
}

func GetError(err error) string {
    str := ""
    if err == nil {
        str = `{ "status": 0, "message": "successful!" }`
    } else {
        str = `{ "status": 1, "message": "unsuccessful!" }`
    }

    return str
}

type BasicSrvController struct {
}

func NewBasicSrvController() *BasicSrvController {
	return &BasicSrvController{}
}

func (controller *BasicSrvController) Query(w http.ResponseWriter, r *http.Request) {
    basic := &model.BasicSrvModel{}
    res, err := basic.FindAll()
    if err != nil {
        res = "[]"
    }

    SendBack(w, res)
}

func (controller *BasicSrvController) Get(w http.ResponseWriter, r *http.Request) {
    id := GetId(r)

    basic := &model.BasicSrvModel{}
    res, err := basic.Find(id)
    if err != nil {
        res = "{}"
    }

    SendBack(w, res)
}

func (controller *BasicSrvController) New(w http.ResponseWriter, r *http.Request) {
    res := "{}"

    //r.ParseForm()
    defer r.Body.Close()
    data, err := ioutil.ReadAll(r.Body)
    if err == nil {
        basic := &model.BasicSrvModel{}
        res, err = basic.Insert(string(data))
        if err != nil {
            res = "{}"
        }
    }

    SendBack(w, res)
}

func (controller *BasicSrvController) Update(w http.ResponseWriter, r *http.Request) {
    res := "{}"
    id := GetId(r)

    //r.ParseForm()
    defer r.Body.Close()
    data, err := ioutil.ReadAll(r.Body)
    if err == nil {
        basic := &model.BasicSrvModel{}
        res, err = basic.Update(id, string(data))
        if err != nil {
            res, err = basic.Find(id)
            if err != nil {
                res = "{}"
            }
        }
    }

    SendBack(w, res)
}

func (controller *BasicSrvController) Delete(w http.ResponseWriter, r *http.Request) {
    id := GetId(r)

    basic := &model.BasicSrvModel{}
    err := basic.Delete(id)
    ret := GetError(err)

    SendBack(w, res)
}

