/**
* @file app.go
* @brief web frame for using angular 
* @author yingx
* @date 2015-12-12
 */

package main

import (
	"net/http"
	"strings"
	"reflect"
	"log"
	"fmt"
	"flag"
	controller "go-angular/server/controller"
)

type params struct {
    host string
    port int
}

func handleCommandLine() *params {
    p := params{}

    flag.StringVar(&p.host, "host", "0.0.0.0", "host to listen to")
    flag.IntVar(&p.port, "port", 9898, "port to listen to")
    flag.Parse()

    return &p
}

type Controller func() reflect.Value

func controllerAction(w http.ResponseWriter, r *http.Request, c Controller) {
	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	action := ""
	if len(parts) > 1 {
		action = parts[1]
	}
	action = strings.Title(action) + "Action"

	controller := c()
	method := controller.MethodByName(action)
	if !method.IsValid() {
		method = controller.MethodByName("IndexAction")
	}
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	method.Call([]reflect.Value{responseValue, requestValue})
}

func controllerResty(w http.ResponseWriter, r *http.Request, c Controller) {
	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	action := ""
	id := ""
	if len(parts) > 1 {
		id = parts[1]
	}
	method := r.Method
	switch method {
		case "GET":
			if id == "" {
				action = "Query"
			} else {
				action = "Get"
			}
		case "POST":
            //-1 represent new item
			if id == "-1" {
				action = "New"
			} else {
				action = "Update"
			}
		case "DELETE":
			action = "Delete"
		case "PUT":
			action = "Update"
		/*	
		case "HEAD":
			action = "Head"
		case "PATCH":
			action = "Patch"
		case "OPTIONS":
			action = "Options"
		*/
		default:
			action = "Query"
	}

	controller := c()
	operation := controller.MethodByName(action)
	if !operation.IsValid() {
		operation = controller.MethodByName("Get")
	}
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	operation.Call([]reflect.Value{responseValue, requestValue})
}

func contentHandler(w http.ResponseWriter, r *http.Request) {
	content := controller.NewContentController()
	controller := reflect.ValueOf(content)
	controllerAction(w, r, func() reflect.Value {
		return controller
		})
}

func recipeHandler(w http.ResponseWriter, r *http.Request) {
	recipe := controller.NewRecipeController()
	controller := reflect.ValueOf(recipe)
	controllerResty(w, r, func() reflect.Value {
		return controller
		})
}

func main() {
    p := handleCommandLine()

	//set static directory	
	http.Handle("/assets/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.Handle("/extensions/", http.FileServer(http.Dir("public")))
	http.Handle("/icons/", http.FileServer(http.Dir("public")))
	http.Handle("/imges/", http.FileServer(http.Dir("public")))
	http.Handle("/js/", http.FileServer(http.Dir("public")))
	//set app directory 
	http.Handle("/app/", http.FileServer(http.Dir("client")))

	http.HandleFunc("/", contentHandler)
	http.HandleFunc("/content/", contentHandler)
	http.HandleFunc("/recipe/", recipeHandler)
    server := fmt.Sprintf("%s:%d", p.host, p.port)
	err := http.ListenAndServe(server, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
