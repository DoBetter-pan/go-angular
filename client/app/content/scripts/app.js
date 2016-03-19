/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

var app = angular.module('guthub', ['ngRoute', 'guthub.services', 'guthub.directives']);

app.config(['$interpolateProvider', function($interpolateProvider){
  $interpolateProvider.startSymbol('[[').endSymbol(']]');
}]);

app.config(['$routeProvider', function($routeProvider){
  $routeProvider.when('/', {
    controller: 'ListCtrl',
    resolve: {
      recipes: function(MultiRecipeLoader) {
        return MultiRecipeLoader();
      }
    },
    templateUrl: 'app/content//views/list.html'
  }).when('edit/:recipeId', {
    controller: 'EditCtrl',
    resolve: {
      recipe: function(RecipeLoader){
        return RecipeLoader();
      }
    },
    templateUrl: 'app/content/views/recipeForm.html'
  }).when('view/:recipeId', {
    controller: 'ViewCtrl',
    resolve: {
      recipe: function(RecipeLoader) {
        return RecipeLoader();
      }
    },
    templateUrl: 'app/content/views/viewRecipe.html'
  }).when('new', {
    controller: 'NewCtrl',
    templateUrl: 'app/content/views/recipeForm.html'
  }).otherwise({redirectTo: '/'});
}]);
