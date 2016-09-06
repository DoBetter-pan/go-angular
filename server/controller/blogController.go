/**
* @file blogController.go
* @brief blog controller
* @author yingx
* @date 2015-02-27
 */

package controller

import (
    "fmt"
	"net/http"
	"log"
	"html/template"
	model "go-angular/server/model"
)

type BlogLink struct {
    Name string
    Url string
}

type BlogMenu struct {
    MainMenu BlogLink
    SubMenu []BlogLink
}

type BlogController struct {
    Shortcuts []BlogLink
    Menus []BlogMenu
}

func NewBlogController() *BlogController {

    controller := &BlogController{}

    linkModel := &model.LinkSrvModel{}
    linkList, _ := linkModel.FindAllLinks()
    controller.Shortcuts = make([]BlogLink, 0, 6)
    for _, link := range(linkList) {
        controller.Shortcuts = append(controller.Shortcuts, BlogLink{Name:link.Name, Url:link.Url})
    }

    menuModel := &model.MenuSrvModel{}
    menuList, _ := menuModel.FindAllMenus()
    controller.Menus = make([]BlogMenu, 0, 12)
    for _, menu := range(menuList) {
        var blogMenu BlogMenu
        blogMenu.MainMenu.Name = menu.MainMenu.Name
        blogMenu.MainMenu.Url = menu.MainMenu.Url
        blogMenu.SubMenu = make([]BlogLink, 0, 12)
        for _, subMenu := range(menu.SubMenu){
            blogMenu.SubMenu = append(blogMenu.SubMenu, BlogLink{Name:subMenu.Name, Url:subMenu.Url})
        }
        controller.Menus = append(controller.Menus, blogMenu)
    }

    fmt.Println("===========>", controller)

    return controller
}


func (controller *BlogController) IndexAction(w http.ResponseWriter, r *http.Request) {
     tmpl, err := template.ParseFiles("client/app/blog/index.html")
    if err != nil {
        log.Fatal("BlogController::IndexAction: ", err)
    }

    err = tmpl.Execute(w, controller)
}
