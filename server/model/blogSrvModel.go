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

type ArticleList struct {
    Page  int64 `json:"page"`
    Total int64 `json:"total"`
    HasPrevious  int64 `json:"previous"`
    PreviousUrl string `json:"previousUrl"`
    HasNext  int64 `json:"next"`
    NextUrl string `json:"nextUrl"`
    Articles []Article `json:"articles"`
}

type ArticleDesc struct {
    Id int64 `json:"id"`
    Title  string `json:"title"`
    TitleHtml  string `json:"titleHtml"`
}

type ArticlesInCategory struct {
    Id int64 `json:"id"`
    Name  string `json:"name"`
    Url  string `json:"url"`
    Description  string `json:"description"`
    Articles []ArticleDesc `json:"articles"`
}

var blogSqls map[string] string = map[string] string {
    "queryindex":"select id, title, titleHtml from ng_blog_article where categoryId=? order by posted desc limit 3",
    "querylist":"select article.id, user.name, title, titleHtml, content, contentHtml, IFNULL(excerpt, '') as excerpt, IFNULL(excerptHtml, '') as excerptHtml, section.name, category.name, commentsCount, status, posted, lastMod, IFNULL(expires, '') as expires from ng_blog_article article, ng_blog_section section, ng_blog_category category, ng_blog_user user where article.sectionId=section.id and categoryId=category.id and authorId=user.id",
    "querycount":"select count(*) as total from ng_blog_article where 1",
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

func (model *BlogSrvModel) FindAllByKeyValue(key string, value, page int64) (string, error) {
    articleList := ArticleList{}
    articleList.Articles = make([]Article, 0, 10)

    condition := ""
    conditionList := ""
    offset := page * 10
    if key == "s" {
        condition = " and sectionId=?"
        conditionList = fmt.Sprintf(" and article.sectionId=? order by posted desc limit 10 offset %d", offset)
    } else if key == "c" {
        condition = " and categoryId=?"
        conditionList = fmt.Sprintf(" and article.categoryId=? order by posted desc limit 10 offset %d", offset)
    }
    sqlCount := blogSqls["querycount"] + condition

    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return "", err
    }

    err = tx.QueryRow(sqlCount, value).Scan(&articleList.Total)
    if err != nil {
        tx.Rollback()
        return "", err
    }
    articleList.Page = page
    if page > 0 {
        articleList.HasPrevious = 1
        articleList.PreviousUrl = fmt.Sprintf("/#/list?%s=%d&p=%d", key, value, page - 1)
    } else {
        articleList.HasPrevious = 0
        articleList.PreviousUrl = ""
    }
    if articleList.Total - (page + 1) * 10 > 0 {
        articleList.HasNext = 1
        articleList.NextUrl = fmt.Sprintf("/#/list?%s=%d&p=%d", key, value, page + 1)
    } else {
        articleList.HasNext = 0
        articleList.NextUrl = ""
    }

    sqlList := blogSqls["querylist"] + conditionList
    rowsArticle, err := tx.Query(sqlList, value)
    if err != nil {
        tx.Rollback()
        return "", err
    }
    defer rowsArticle.Close()

    for rowsArticle.Next() {
        var art Article

        err = rowsArticle.Scan(&art.Id, &art.Author, &art.Title, &art.TitleHtml, &art.Content, &art.ContentHtml, &art.Excerpt, &art.ExcerptHtml, &art.Section, &art.Category, &art.CommentCount, &art.Status, &art.Posted, &art.LastMod, &art.Expires)
        if err == nil {
            articleList.Articles = append(articleList.Articles, art)
        }
    }

    //check error
    if err = rowsArticle.Err(); err != nil {
        tx.Rollback()
        return "", err
    }

    tx.Commit()

    //change into array
    articleListArray := make([]ArticleList, 0, 1)
    articleListArray = append(articleListArray, articleList)
    data, err := json.Marshal(articleListArray)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *BlogSrvModel) FindAll() (string, error) {
    artsInCatList := make([]ArticlesInCategory, 0, 20)

    sql := blogSqls["category"] + " and isPage=1"
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
        var artsInCat ArticlesInCategory
        artsInCat.Articles = make([]ArticleDesc, 0, 5)
        err = rows.Scan(&artsInCat.Id, &artsInCat.Name, &artsInCat.Url, &artsInCat.Description)
        if err == nil {
            artsInCatList = append(artsInCatList, artsInCat)
        }
    }

    //check error
    if err = rows.Err(); err != nil {
        tx.Rollback()
        return "", err
    }

    artsInCatLen := len(artsInCatList)
    for i := 0; i < artsInCatLen; i++ {
        rowsArticle, err := tx.Query(blogSqls["queryindex"], artsInCatList[i].Id)
        if err != nil {
            tx.Rollback()
            return "", err
        }
        defer rowsArticle.Close()

        for rowsArticle.Next() {
            var art ArticleDesc

            err = rowsArticle.Scan(&art.Id, &art.Title, &art.TitleHtml )
            if err == nil {
                artsInCatList[i].Articles = append(artsInCatList[i].Articles, art)
            }
        }

        //check error
        if err = rows.Err(); err != nil {
            tx.Rollback()
            return "", err
        }
    }
    tx.Commit()

    data, err := json.Marshal(artsInCatList)
    if err != nil {
        return "", err
    }

    return string(data), nil
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
