/**
* @file blogSrvModel.go
* @brief blog service model
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

type Blog struct {
     Id int64 `json:"id"`
     Name  string `json:"name"`
     Age int64 `json:"age"`
     Sex int64 `json:"sex"`
}

var blogSqls map[string] string = map[string] string {
    "query":"select id, name, age, sex from blog",
    "queryone":"select id, name, age, sex from blog where id=?",
    "insert":"insert into blog( name, age, sex) values( ?, ?, ?)",
    "update":"update blog set name=?, age=?, sex=? where id=?",
    "delete":"delete from blog where id=?",
}

type BlogSrvModel struct {
}

func NewBlogSrvModel() *BlogSrvModel {
	return &BlogSrvModel{}
}

func (model *BlogSrvModel) FindAll() (string, error) {
    dbconnection := dbwrapper.GetDatabaseConnection()
    rows, err := dbconnection.DB.Query(blogSqls["query"])
    if err != nil {
        return "", err
    }
    defer rows.Close()

    blogList := make([]Blog, 0, 10)
    for rows.Next() {
        var blog Blog
        err = rows.Scan( &blog.Id, &blog.Name, &blog.Age, &blog.Sex)
        if err == nil {
            blogList = append(blogList, blog)
        }
    }

    //check error
    if err = rows.Err(); err != nil {
        return "", err
    }

    data, err := json.Marshal(blogList)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *BlogSrvModel) Find(id int64) (string, error) {
    var blog Blog
    dbconnection := dbwrapper.GetDatabaseConnection()
    err := dbconnection.DB.QueryRow(blogSqls["queryone"], id).Scan( &blog.Id, &blog.Name, &blog.Age, &blog.Sex)
    if err != nil {
        return "", err
    }

    data, err := json.Marshal(blog)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *BlogSrvModel) Insert(str string) (string, error) {
    var blog Blog

    err := json.Unmarshal([]byte(str), &blog)
    if err != nil {
        return "", err
    }

    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return "", err
    }

    res, err := tx.Exec(blogSqls["insert"],  blog.Name, blog.Age, blog.Sex, blog.Id)
    if err != nil {
        tx.Rollback()
        return "", err
    }

    blogid, err := res.LastInsertId()
    if err != nil {
        tx.Rollback()
        return "", err
    }

    tx.Commit()

    blog.Id = blogid
    data, err := json.Marshal(blog)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *BlogSrvModel) Update(id int64, str string) (string, error) {
    var blog Blog

    err := json.Unmarshal([]byte(str), &blog)
    if err != nil {
        return "", err
    }

    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return "", err
    }

    //just update, not check if it is same before updating. It may be supported in future
    _, err = tx.Exec(blogSqls["update"],  blog.Name, blog.Age, blog.Sex, blog.Id)
    if err != nil {
        tx.Rollback()
        return "", err
    }

    tx.Commit()

    return str, nil
}

func (model *BlogSrvModel) Delete(id int64) error {
    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return err
    }

    _, err = tx.Exec(blogSqls["delete"], id)
    if err != nil {
        tx.Rollback()
        return err
    }

    tx.Commit()

    return nil
}
