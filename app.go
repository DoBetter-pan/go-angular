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
	"strconv"
	"reflect"
	"log"
	"fmt"
	"flag"
	controller "go-angular/server/controller"
	session "go-angular/server/session"
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

func checkRules(w http.ResponseWriter, r *http.Request, c Controller, action string) bool {
    validated := true
	controllerInstance := c()

    checkRules := controllerInstance.MethodByName("CheckRules")
	if checkRules.IsValid() {
        //get rules
        rulesMapRef := checkRules.Call([]reflect.Value{})

        rulesMap := rulesMapRef[0].Interface().(map[string] []string)
        rules, ok := rulesMap[action]
        if ok {
            //valid, id, name, nonce := session.ValidateSessionByCookie(r)
            valid, id, _, _ := session.ValidateSessionByCookie(r)
            idStr := strconv.FormatInt(id, 10)
            isExit := false
            for _, rule := range(rules) {
                strArray := strings.Split(rule, " ")
                if len(strArray) == 2 {
                    //compute allow, not validated will be deny
                    if strArray[1] == "*" {
                        isExit = true
                        validated = true
                    } else {
                        idArray := strings.Split(strArray[1], ",")
                        for _, v := range(idArray){
                            if v == idStr && valid {
                                isExit = true
                                validated = true
                                break
                            }
                        }
                    }
                    if strArray[0] == "deny" {
                        validated = !validated
                    }
                    if isExit {
                        break
                    }
                }
            }
        }
	}

    return validated
}

func controllerAction(w http.ResponseWriter, r *http.Request, c Controller) {
	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	action := ""
	if len(parts) > 1 {
		action = parts[1]
	}
	action = strings.Title(action) + "Action"

	controllerInstance := c()
	method := controllerInstance.MethodByName(action)
	if !method.IsValid() {
        action = "IndexAction"
		method = controllerInstance.MethodByName(action)
	}

    validated := checkRules(w, r, c, action)
    if validated {
        requestValue := reflect.ValueOf(r)
        responseValue := reflect.ValueOf(w)
        method.Call([]reflect.Value{responseValue, requestValue})
    } else {
        login := controller.NewLoginController()
        login.IndexAction(w, r)
    }
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

	controllerInstance := c()
	operation := controllerInstance.MethodByName(action)
	if !operation.IsValid() {
		operation = controllerInstance.MethodByName("Get")
	}
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	operation.Call([]reflect.Value{responseValue, requestValue})
}

func recipeHandler(w http.ResponseWriter, r *http.Request) {
	recipe := controller.NewRecipeController()
	controller := reflect.ValueOf(recipe)
	controllerAction(w, r, func() reflect.Value {
		return controller
		})
}

func recipeSrvHandler(w http.ResponseWriter, r *http.Request) {
	recipeSrv := controller.NewRecipeSrvController()
	controller := reflect.ValueOf(recipeSrv)
	controllerResty(w, r, func() reflect.Value {
		return controller
		})
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
    blog := controller.NewBlogController()
    controller := reflect.ValueOf(blog)
    controllerAction(w, r, func() reflect.Value {
        return controller
        })
}

func blogSrvHandler(w http.ResponseWriter, r *http.Request) {
    blogSrv := controller.NewBlogSrvController()
    controller := reflect.ValueOf(blogSrv)
    controllerResty(w, r, func() reflect.Value {
        return controller
        })
}

func linkHandler(w http.ResponseWriter, r *http.Request) {
    link := controller.NewLinkController()
    controller := reflect.ValueOf(link)
    controllerAction(w, r, func() reflect.Value {
        return controller
        })
}

func linkSrvHandler(w http.ResponseWriter, r *http.Request) {
    linkSrv := controller.NewLinkSrvController()
    controller := reflect.ValueOf(linkSrv)
    controllerResty(w, r, func() reflect.Value {
        return controller
        })
}

func menuHandler(w http.ResponseWriter, r *http.Request) {
    menu := controller.NewMenuController()
    controller := reflect.ValueOf(menu)
    controllerAction(w, r, func() reflect.Value {
        return controller
        })
}

func menuSrvHandler(w http.ResponseWriter, r *http.Request) {
    menuSrv := controller.NewMenuSrvController()
    controller := reflect.ValueOf(menuSrv)
    controllerResty(w, r, func() reflect.Value {
        return controller
        })
}

func sectionHandler(w http.ResponseWriter, r *http.Request) {
    section := controller.NewSectionController()
    controller := reflect.ValueOf(section)
    controllerAction(w, r, func() reflect.Value {
        return controller
        })
}

func sectionSrvHandler(w http.ResponseWriter, r *http.Request) {
    sectionSrv := controller.NewSectionSrvController()
    controller := reflect.ValueOf(sectionSrv)
    controllerResty(w, r, func() reflect.Value {
        return controller
        })
}

func categoryHandler(w http.ResponseWriter, r *http.Request) {
    category := controller.NewCategoryController()
    controller := reflect.ValueOf(category)
    controllerAction(w, r, func() reflect.Value {
        return controller
        })
}

func categorySrvHandler(w http.ResponseWriter, r *http.Request) {
    categorySrv := controller.NewCategorySrvController()
    controller := reflect.ValueOf(categorySrv)
    controllerResty(w, r, func() reflect.Value {
        return controller
        })
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    login := controller.NewLoginController()
    controller := reflect.ValueOf(login)
    controllerAction(w, r, func() reflect.Value {
        return controller
        })
}

func loginSrvHandler(w http.ResponseWriter, r *http.Request) {
    loginSrv := controller.NewLoginSrvController()
    controller := reflect.ValueOf(loginSrv)
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

	//http.HandleFunc("/", recipeHandler)
	http.HandleFunc("/recipe", recipeHandler)
	http.HandleFunc("/recipe/", recipeHandler)	
	http.HandleFunc("/recipesrv", recipeSrvHandler)
	http.HandleFunc("/recipesrv/", recipeSrvHandler)	
    http.HandleFunc("/blog", blogHandler)
    http.HandleFunc("/blog/", blogHandler)    
    http.HandleFunc("/blogsrv", blogSrvHandler)
    http.HandleFunc("/blogsrv/", blogSrvHandler)    
    http.HandleFunc("/link", linkHandler)
    http.HandleFunc("/link/", linkHandler)    
    http.HandleFunc("/linksrv", linkSrvHandler)
    http.HandleFunc("/linksrv/", linkSrvHandler)    
    http.HandleFunc("/menu", menuHandler)
    http.HandleFunc("/menu/", menuHandler)    
    http.HandleFunc("/menusrv", menuSrvHandler)
    http.HandleFunc("/menusrv/", menuSrvHandler)    
    http.HandleFunc("/section", sectionHandler)
    http.HandleFunc("/section/", sectionHandler)
    http.HandleFunc("/sectionsrv", sectionSrvHandler)
    http.HandleFunc("/sectionsrv/", sectionSrvHandler)
    http.HandleFunc("/category", categoryHandler)
    http.HandleFunc("/category/", categoryHandler)
    http.HandleFunc("/categorysrv", categorySrvHandler)
    http.HandleFunc("/categorysrv/", categorySrvHandler)
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/login/", loginHandler)    
    http.HandleFunc("/loginsrv", loginSrvHandler)
    http.HandleFunc("/loginsrv/", loginSrvHandler)    
    server := fmt.Sprintf("%s:%d", p.host, p.port)
	err := http.ListenAndServe(server, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
