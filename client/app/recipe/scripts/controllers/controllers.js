/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

app.controller('ListCtrl', ['$scope', 'recipes', function($scope, recipes){
    //console.log(recipes);
    $scope.recipes = recipes;
}]);

app.controller('ViewCtrl', ['$scope', '$location', 'recipe', function($scope, $location, recipe){
    $scope.recipe = recipe;

    $scope.edit = function(){
        $location.path('/edit/' + recipe.id);
    }
}]);

app.controller('EditCtrl', ['$scope', '$location', 'recipe', function($scope, $location, recipe){
    $scope.recipe = recipe;

    $scope.save = function(){
        $scope.recipe.$save(function(recipe){
            $location.path('/view/' + recipe.id);
        });
    };

    $scope.remove = function(){
        $scope.recipe.$remove(function(recipe){
            $location.path('/');
        });
    };
}]);

app.controller('NewCtrl', ['$scope', '$location', 'RecipeSrv', function($scope, $location, RecipeSrv){
    $scope.recipe = new RecipeSrv({
        id: -1,
        ingredients: [{}]
    });

    $scope.save = function(){
        $scope.recipe.$save(function(recipe){
            $location.path('/view/' + recipe.id);
        });
    };
}]);

app.controller('IngredientsCtrl', ['$scope', function($scope){
    $scope.addIngredient = function(){
        var ingredients = $scope.recipe.ingredients;
        ingredients[ingredients.length] = {};
    };
    $scope.removeIngredient = function(index){
        $scope.recipe.ingredients.splice(index, 1);
    }
}])

