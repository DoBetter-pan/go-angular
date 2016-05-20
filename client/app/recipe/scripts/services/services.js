/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

var services = angular.module('guthub.services', ['ngResource']);

/*get, save, query, remove, delete*/
services.factory('RecipeSrv', ['$resource', function($resource){
  return $resource('/recipesrv/:id', {id: '@id'});
}]);

services.factory('MultiRecipeLoader', ['RecipeSrv', '$q', function(RecipeSrv, $q){
  return function() {
    var delay = $q.defer();
    RecipeSrv.query(function(recipes){
      delay.resolve(recipes);
    }, function(){
      delay.reject('Unable to fetch recipes');
    });
    return delay.promise;
  }
}]);

services.factory('RecipeLoader', ['RecipeSrv', '$route', '$q', function(RecipeSrv, $route, $q){
  return function() {
    var delay = $q.defer();
    RecipeSrv.get({id:$route.current.params.recipeId}, function(recipe){
      delay.resolve(recipe);
    }, function(){
      delay.reject('Unable to fetch recipe ' + $route.current.params.recipeId);
    });
    return delay.promise;
  }
}]);

