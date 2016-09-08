/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

var app = angular.module('blog', ['ngRoute', 'blog.services', 'blog.directives']);

app.config(['$interpolateProvider', function($interpolateProvider){
    $interpolateProvider.startSymbol('[[').endSymbol(']]');
}]);

app.config(['$routeProvider', function($routeProvider){
    $routeProvider.when('/', {
        controller: 'IndexCtrl',
        resolve: {
            articles: function(MultiBlogLoader) {
                return MultiBlogLoader();
            }
        },
        templateUrl: '/app/blog/views/index.html'
    }).when('/list', {
        controller: 'ListCtrl',
        resolve: {
            articles: function(MultiBlogLoader) {
                return MultiBlogLoader();
            }
        },
        templateUrl: '/app/blog/views/list.html'
    }).when('/edit/:blogId', {
        controller: 'EditCtrl',
        resolve: {
            article: function(BlogLoader){
                return BlogLoader();
            }
        },
        templateUrl: '/app/blog/views/form.html'
    }).when('/view/:blogId', {
        controller: 'ViewCtrl',
        resolve: {
            article: function(BlogLoader) {
                return BlogLoader();
            }
        },
        templateUrl: '/app/blog/views/view.html'
    }).when('/new', {
        controller: 'NewCtrl',
        templateUrl: '/app/blog/views/form.html'
    }).otherwise({redirectTo: '/'});
}]);
