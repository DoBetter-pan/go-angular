/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

var directives = angular.module('guthub.directives', []);

directives.directive('butterbar', ['$rootScope', function($rootScope){
  return {
    link: function(scope, element, attrs) {
      element.addClass('hide');

      $rootScope.$on('$routeChangeStart', function(){
        element.removeClass('hide');
      });

      $rootScope.$on('$routeChangeSuccess', function(){
        element.addClass('hide');
      });
    }
  };
}]);

directives.directive('focus', function(){
  return {
    link: function(scope, element, attrs){
      element[0].focus();
    }
  };
});
