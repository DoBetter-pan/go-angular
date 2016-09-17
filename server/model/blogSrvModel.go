/**
* @file blogSrvModel.go
* @brief blog service model
* @author yingx
* @date 2015-02-27
 */

package model

import (
	"fmt"
    "encoding/json"
    _ "database/sql"
    _ "github.com/go-sql-driver/mysql"
	dbwrapper "go-angular/server/datawrapper"
)

type Article struct {
    Id int64 `json:"id"`
    Author  string `json:"author"`
    Title  string `json:"title"`
    TitleHtml  string `json:"titleHtml"` 
    Content  string `json:"content"` 
    ContentHtml  string `json:"contentHtml"`
    Excerpt  string `json:"excerpt"` 
    ExcerptHtml  string `json:"excerptHtml"`
    Section string `json:"section"`
    Category string `json:"category"`   
    CommentCount int64 `json:"commentCount"` 
    Status int64 `json:"status"`                 
    Posted  string `json:"posted"`                          
    LastMod string `json:"lastMod"`
    Expires string `json:"expires"`
}

type CategoryArticle struct {
    Id int64 `json:"id"`
    Name  string `json:"name"`
    Url  string `json:"url"`
    Description  string `json:"description"`
    Articles []Article `json:"articles"`              
}

var blogSqls map[string] string = map[string] string {
    "query":"select article.id, user.name, title, titleHtml, content, contentHtml, IFNULL(excerpt, '') as excerpt, IFNULL(excerptHtml, '') as excerptHtml, section.name, category.name, commentsCount, status, posted, lastMod, IFNULL(expires, '') as expires from ng_blog_article article, ng_blog_section section, ng_blog_category category, ng_blog_user user where article.sectionId=section.id and categoryId=category.id and authorId=user.id and article.sectionId=?",
    "queryone":"select id, name, age, sex from blog where id=?",
    "insert":"insert into blog( name, age, sex) values( ?, ?, ?)",
    "update":"update blog set name=?, age=?, sex=? where id=?",
    "delete":"delete from blog where id=?",
    "category":"select id, name, url, description from ng_blog_category where 1", 
}

type BlogSrvModel struct {
}

func NewBlogSrvModel() *BlogSrvModel {
	return &BlogSrvModel{}
}

func (model *BlogSrvModel) FindAllByKeyValue(key, value string) (string, error) {
    catList := make([]CategoryArticle, 0, 20)

    condition := " and isPage=1"
    if key != "" && value != "" {
        if key == "s" {
            condition = fmt.Sprintf(" and sectionId=%s", value)
        } else if key == "c" {
            condition = fmt.Sprintf(" and id=%s", value)                        
        }
    }

    sql := blogSqls["category"] + condition
    fmt.Println("==========>", sql)

    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return "", err
    }

    rows, err := tx.Query(sql)
    if err != nil {
        tx.Rollback()
        return "", err
    }
    defer rows.Close()

    for rows.Next() {
        var cat CategoryArticle
        cat.Articles = make([]Article, 0, 20)
        err = rows.Scan( &cat.Id, &cat.Name, &cat.Url, &cat.Description)
        if err == nil {
            catList = append(catList, cat)
        }
    }

    //check error
    if err = rows.Err(); err != nil {
        tx.Rollback()
        return "", err
    }

    catLen := len(catList)
    for i := 0; i < catLen; i++ {
        rowsArticle, err := tx.Query(blogSqls["query"], catList[i].Id)
        if err != nil {
            tx.Rollback()
            return "", err
        }
        defer rowsArticle.Close()

        for rowsArticle.Next() {
            var art Article
            
            err = rowsArticle.Scan(&art.Id, &art.Author, &art.Title, &art.TitleHtml, &art.Content, &art.ContentHtml, &art.Excerpt, &art.ExcerptHtml, &art.Section, &art.Category, &art.CommentCount, &art.Status, &art.Posted, &art.LastMod, &art.Expires)
            if err == nil {
                catList[i].Articles = append(catList[i].Articles, art)
            }
            fmt.Println("----------->", err) 
        }

        //check error
        if err = rows.Err(); err != nil {
            tx.Rollback()
            return "", err
        }
    }
    tx.Commit()

    data, err := json.Marshal(catList)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *BlogSrvModel) FindAll() (string, error) {
    /*
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
    */
    return "", nil
}

func (model *BlogSrvModel) Find(id int64) (string, error) {
    /*
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
    */
    return "", nil    
}

func (model *BlogSrvModel) Insert(str string) (string, error) {
    /*
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
    */
    return "", nil    
}

func (model *BlogSrvModel) Update(id int64, str string) (string, error) {
    /*
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
    */
    return "", nil    
}

func (model *BlogSrvModel) Delete(id int64) error {
    /*
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
    */

    return nil
}
