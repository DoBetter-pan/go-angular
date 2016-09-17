/**
* @file menuSrvModel.go
* @brief menu service model
* @author yingx
* @date 2015-02-27
 */

package model

import (
	_ "fmt"
    "encoding/json"
    _ "database/sql"
    _ "github.com/go-sql-driver/mysql"
	dbwrapper "go-angular/server/datawrapper"
)

type Section struct {
     Id int64 `json:"id"`
     Name  string `json:"name"`
     Url  string `json:"url"`
     Authority int64 `json:"authority"`
}

type Category struct {
     Id int64 `json:"id"`
     Name  string `json:"name"`
     Url  string `json:"url"`
}

type Menu struct {
     MainMenu Section
     SubMenu []Category
}

var sectionSqls map[string] string = map[string] string {
    "query":"select id, name, url, authority from ng_blog_section",
    "queryone":"select id, name, url, authority from ng_blog_section where id=?",
    "insert":"insert into ng_blog_section( name, url, authority) values( ?, ?, ?)",
    "update":"update ng_blog_section set name=?, url=?, authority=? where id=?",
    "delete":"delete from ng_blog_section where id=?",
}

var categorySqls map[string] string = map[string] string {
    "query":"select id, name, url from ng_blog_category where sectionId=?",
    "queryone":"select id, name, url from ng_blog_category where id=?",
    "insert":"insert into ng_blog_category( name, url, sectionId) values( ?, ?, ?)",
    "update":"update ng_blog_category set name=?, url=?, sectionId=? where id=?",
    "delete":"delete from ng_blog_category where id=?",
}

type MenuSrvModel struct {
}

func NewMenuSrvModel() *MenuSrvModel {
	return &MenuSrvModel{}
}

func (model *MenuSrvModel) FindAllMenus() ([]Menu, error) {
    menuList := make([]Menu, 0, 10)
    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return menuList, err
    }

    rows, err := tx.Query(sectionSqls["query"])
    if err != nil {
        tx.Rollback()
        return menuList, err
    }
    defer rows.Close()

    for rows.Next() {
        var menu Menu
        menu.SubMenu = make([]Category, 0, 10)
        err = rows.Scan( &menu.MainMenu.Id, &menu.MainMenu.Name, &menu.MainMenu.Url, &menu.MainMenu.Authority)
        if err == nil {
            menuList = append(menuList, menu)
        }
    }

    //check error
    if err = rows.Err(); err != nil {
        tx.Rollback()
        return menuList, err
    }

    menuLen := len(menuList)
    for i := 0; i < menuLen; i++ {
        rowsCat, err := tx.Query(categorySqls["query"], menuList[i].MainMenu.Id)
        if err != nil {
            tx.Rollback()
            return menuList, err
        }
        defer rowsCat.Close()

        for rowsCat.Next() {
            var cat Category
            err = rowsCat.Scan( &cat.Id, &cat.Name, &cat.Url)
            if err == nil {
                menuList[i].SubMenu = append(menuList[i].SubMenu, cat)
            } 
        }

        //check error
        if err = rows.Err(); err != nil {
            tx.Rollback()
            return menuList, err
        }
    }

    tx.Commit()

    return menuList, nil
}

func (model *MenuSrvModel) FindAll() (string, error) {
    menuList, err := model.FindAllMenus()
    if err != nil {
        return "", err
    }

    data, err := json.Marshal(menuList)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *MenuSrvModel) Find(id int64) (string, error) {
    /*
    var menu Menu
    dbconnection := dbwrapper.GetDatabaseConnection()
    err := dbconnection.DB.QueryRow(menuSqls["queryone"], id).Scan( &menu.Id, &menu.Name, &menu.Age, &menu.Sex)
    if err != nil {
        return "", err
    }

    data, err := json.Marshal(menu)
    if err != nil {
        return "", err
    }

    return string(data), nil
    */
    return "", nil
}

func (model *MenuSrvModel) Insert(str string) (string, error) {
    /*
    var menu Menu

    err := json.Unmarshal([]byte(str), &menu)
    if err != nil {
        return "", err
    }

    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return "", err
    }

    res, err := tx.Exec(menuSqls["insert"],  menu.Name, menu.Age, menu.Sex, menu.Id)
    if err != nil {
        tx.Rollback()
        return "", err
    }

    menuid, err := res.LastInsertId()
    if err != nil {
        tx.Rollback()
        return "", err
    }

    tx.Commit()

    menu.Id = menuid
    data, err := json.Marshal(menu)
    if err != nil {
        return "", err
    }

    return string(data), nil
    */
    return "", nil
}

func (model *MenuSrvModel) Update(id int64, str string) (string, error) {
    /*
    var menu Menu

    err := json.Unmarshal([]byte(str), &menu)
    if err != nil {
        return "", err
    }

    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return "", err
    }

    //just update, not check if it is same before updating. It may be supported in future
    _, err = tx.Exec(menuSqls["update"],  menu.Name, menu.Age, menu.Sex, menu.Id)
    if err != nil {
        tx.Rollback()
        return "", err
    }

    tx.Commit()

    return str, nil
    */
    return "", nil
}

func (model *MenuSrvModel) Delete(id int64) error {
    /*
    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return err
    }

    _, err = tx.Exec(menuSqls["delete"], id)
    if err != nil {
        tx.Rollback()
        return err
    }

    tx.Commit()
    */

    return nil
}
