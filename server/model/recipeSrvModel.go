/**
* @file recipeSrvModel.go
* @brief recipe service model
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

type Ingredient struct {
    Amount int32                 `json:"amount"`
    AmountUnits string           `json:"amountUnits"`
    IngredientName string        `json:"ingredientName"`
}

type Recipe struct {
    Id int64                 `json:"id"`
    Title string             `json:"title"`
    Description string       `json:"description"`
    Instructions string      `json:"instructions"`
    Ingredients []Ingredient `json:"ingredients"`
}

type RecipeSrvModel struct {
}

func NewRecipeSrvModel() *RecipeSrvModel {
	return &RecipeSrvModel{}
}

func (model *RecipeSrvModel) FindAll() (string, error) {
    dbconnection := dbwrapper.GetDatabaseConnection()
    rows, err := dbconnection.DB.Query("select id, title, description, instructions from recipe")
    if err != nil {
        return "", err
    }
    defer rows.Close()

    recipeList := make([]Recipe, 0, 10)
    for rows.Next() {
        var recipe Recipe
        err = rows.Scan(&recipe.Id, &recipe.Title, &recipe.Description, &recipe.Instructions)
        if err == nil {
            recipe.Ingredients = make([]Ingredient, 0, 10)
            recipeList = append(recipeList, recipe)
        }
    }

    //check error
    if err = rows.Err(); err != nil {
        return "", err
    }

    //deal with ingredients
    stmt, err := dbconnection.DB.Prepare("select name, unit, amount from ingredient, formula where ingredient.id = ingredientid and recipeid=?")
    if err == nil {
        defer stmt.Close()
        recipeCount := len(recipeList)
        for i := 0; i < recipeCount; i++ {
            if rows, err = stmt.Query(recipeList[i].Id); err == nil {
                for rows.Next() {
                    var ingredient Ingredient
                    if err = rows.Scan(&ingredient.IngredientName, &ingredient.AmountUnits, &ingredient.Amount); err == nil {
                        recipeList[i].Ingredients = append(recipeList[i].Ingredients, ingredient)
                    }
                }
            }
        }
    }

    data, err := json.Marshal(recipeList)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *RecipeSrvModel) Find(id int64) (string, error) {
    var recipe Recipe
    recipe.Ingredients = make([]Ingredient, 0, 10)
    dbconnection := dbwrapper.GetDatabaseConnection()
    err := dbconnection.DB.QueryRow("select id, title, description, instructions from recipe where id=?", id).Scan(&recipe.Id, &recipe.Title, &recipe.Description, &recipe.Instructions)
    if err != nil {
        return "", err
    }

    //deal with ingredients
    stmt, err := dbconnection.DB.Prepare("select name, unit, amount from ingredient, formula where ingredient.id = ingredientid and recipeid=?")
    if err == nil {
        defer stmt.Close()
        if rows, err := stmt.Query(recipe.Id); err == nil {
            for rows.Next() {
                var ingredient Ingredient
                if err = rows.Scan(&ingredient.IngredientName, &ingredient.AmountUnits, &ingredient.Amount); err == nil {
                    recipe.Ingredients = append(recipe.Ingredients, ingredient)
                }
            }
        }
    }

    data, err := json.Marshal(recipe)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *RecipeSrvModel) Insert(str string) (string, error) {
    var recipe Recipe

    err := json.Unmarshal([]byte(str), &recipe)
    if err != nil {
        return "", err
    }

    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return "", err
    }

    res, err := tx.Exec("insert into recipe(title,description,instructions) values(?,?,?)", recipe.Title, recipe.Description, recipe.Instructions)
    if err != nil {
        tx.Rollback()
        return "", err
    }

    recipeid, err := res.LastInsertId()
    if err != nil {
        tx.Rollback()
        return "", err
    }

    var ingredientid int64 = 0
    for _, ingredient := range(recipe.Ingredients) {
        row := tx.QueryRow("select id from ingredient where name=? and unit=?", ingredient.IngredientName, ingredient.AmountUnits)
        err = row.Scan(&ingredientid)
        if err != nil {
            res, err = tx.Exec("insert into ingredient(name,unit) values(?,?)", ingredient.IngredientName, ingredient.AmountUnits)
            if err != nil {
                tx.Rollback()
                return "", err
            }
            ingredientid, err = res.LastInsertId()
            if err != nil {
                tx.Rollback()
                return "", err
            }
        }
        _, err = tx.Exec("insert into formula(recipeid,ingredientid,amount) values(?,?,?)", recipeid, ingredientid, ingredient.Amount)
        if err != nil {
            tx.Rollback()
            return "", err
        }
    }
    tx.Commit()

    recipe.Id = recipeid
    data, err := json.Marshal(recipe)
    if err != nil {
        return "", err
    }

    return string(data), nil
}

func (model *RecipeSrvModel) Update(id int64, str string) (string, error) {
    var recipe Recipe

    err := json.Unmarshal([]byte(str), &recipe)
    if err != nil {
        return "", err
    }

    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return "", err
    }

    //just update, not check if it is same before updating. It may be supported in future
    res, err := tx.Exec("update recipe set title=?, description=?, instructions=? where id=?", recipe.Title, recipe.Description, recipe.Instructions, recipe.Id)
    if err != nil {
        tx.Rollback()
        return "", err
    }
    //delete all formula with this recipe
    res, err = tx.Exec("delete from formula where recipeid=?", recipe.Id)
    if err != nil {
        tx.Rollback()
        return "", err
    }

    var ingredientid int64 = 0
    for _, ingredient := range(recipe.Ingredients) {
        row := tx.QueryRow("select id from ingredient where name=? and unit=?", ingredient.IngredientName, ingredient.AmountUnits)
        err = row.Scan(&ingredientid)
        if err != nil {
            res, err = tx.Exec("insert into ingredient(name,unit) values(?,?)", ingredient.IngredientName, ingredient.AmountUnits)
            if err != nil {
                tx.Rollback()
                return "", err
            }
            ingredientid, err = res.LastInsertId()
            if err != nil {
                tx.Rollback()
                return "", err
            }
        }
        _, err = tx.Exec("insert into formula(recipeid,ingredientid,amount) values(?,?,?)", recipe.Id, ingredientid, ingredient.Amount)
        if err != nil {
            tx.Rollback()
            return "", err
        }
    }
    tx.Commit()

    return str, nil
}

func (model *RecipeSrvModel) Delete(id int64) error {
    dbconnection := dbwrapper.GetDatabaseConnection()
    tx, err := dbconnection.DB.Begin()
    if err != nil {
        return err
    }

    _, err = tx.Exec("delete from recipe where id=?", id)
    if err != nil {
        tx.Rollback()
        return err
    }
    //delete all formula with this recipe
    _, err = tx.Exec("delete from formula where recipeid=?", id)
    if err != nil {
        tx.Rollback()
        return err
    }
    tx.Commit()

    return nil
}
