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
        if err := rows.Scan(&recipe.Id, &recipe.Title, &recipe.Description, &recipe.Instructions); err == nil {
            recipe.Ingredients = make([]Ingredient, 0, 10)
            recipeList = append(recipeList, recipe)
        }
    }

    //check error
    if err := rows.Err(); err != nil {
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

