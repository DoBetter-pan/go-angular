/**
* @file loginSrvModel.go
* @brief login service model
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

type Login struct {
     Id int64 `json:"id"`
     Name  string `json:"name"`
     Age int64 `json:"age"`
     Sex int64 `json:"sex"`
}

var loginSqls map[string] string = map[string] string {
    "query":"select id, name, age, sex from login",
    "queryone":"select id, name, age, sex from login where id=?",
    "insert":"insert into login( name, age, sex) values( ?, ?, ?)",
    "update":"update login set name=?, age=?, sex=? where id=?",
    "delete":"delete from login where id=?",
}

type LoginSrvModel struct {
}

func NewLoginSrvModel() *LoginSrvModel {
	return &LoginSrvModel{}
}

func (model *LoginSrvModel) FindAll() (string, error) {
    dbconnection := dbwrapper.GetDatabaseConnection()
    rows, err := dbconnection.DB.Query(loginSqls["query"])
    if err != nil {
        return "", err
    }
    defer rows.Close()

    loginList := make([]Login, 0, 10)
    for rows.Next() {
        var login Login
        err = rows.Scan( &login.Id, &login.Name, &login.Age, &login.Sex)
        if err == nil {
            loginList = append(loginList, login)
        }
    }

    //check error
    if err = rows.Err(); err != nil {
        return "", err
    }

    data, err := json.Marshal(loginList)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *LoginSrvModel) Find(id int64) (string, error) {
    var login Login
    dbconnection := dbwrapper.GetDatabaseConnection()
    err := dbconnection.DB.QueryRow(loginSqls["queryone"], id).Scan( &login.Id, &login.Name, &login.Age, &login.Sex)
    if err != nil {
        return "", err
    }

    data, err := json.Marshal(login)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *LoginSrvModel) Insert(str string) (string, error) {
    var login Login

    err := json.Unmarshal([]byte(str), &login)
    if err != nil {
        return "", err
    }

    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return "", err
    }

    res, err := tx.Exec(loginSqls["insert"],  login.Name, login.Age, login.Sex, login.Id)
    if err != nil {
        tx.Rollback()
        return "", err
    }

    loginid, err := res.LastInsertId()
    if err != nil {
        tx.Rollback()
        return "", err
    }

    tx.Commit()

    login.Id = loginid
    data, err := json.Marshal(login)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *LoginSrvModel) Update(id int64, str string) (string, error) {
    var login Login

    err := json.Unmarshal([]byte(str), &login)
    if err != nil {
        return "", err
    }

    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return "", err
    }

    //just update, not check if it is same before updating. It may be supported in future
    _, err = tx.Exec(loginSqls["update"],  login.Name, login.Age, login.Sex, login.Id)
    if err != nil {
        tx.Rollback()
        return "", err
    }

    tx.Commit()

    return str, nil
}

func (model *LoginSrvModel) Delete(id int64) error {
    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return err
    }

    _, err = tx.Exec(loginSqls["delete"], id)
    if err != nil {
        tx.Rollback()
        return err
    }

    tx.Commit()

    return nil
}
