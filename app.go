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
	controller "go-angular/controller"
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

func angularHandler(w http.ResponseWriter, r *http.Request) {
	angular := controller.NewAngularController()
	controller := reflect.ValueOf(angular)
	controllerAction(w, r, func() reflect.Value {
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

	http.HandleFunc("/", angularHandler)
	http.HandleFunc("/angular/", angularHandler)
    server := fmt.Sprintf("%s:%d", p.host, p.port)
	err := http.ListenAndServe(server, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
