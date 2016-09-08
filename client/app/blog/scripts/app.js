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
    $routeProvider.when('/:object', {
        controller: 'ListCtrl',
        resolve: {
            blogs: function(MultiBlogLoader) {
                return MultiBlogLoader();
            }
        },
        templateUrl: '/app/blog/views/list.html'
    }).when('/:object/:id', {
        controller: 'ListCtrl',
        resolve: {
            blogs: function(MultiBlogLoader) {
                return MultiBlogLoader();
            }
        },
        templateUrl: '/app/blog/views/list.html'
    }).when('/edit/:blogId', {
        controller: 'EditCtrl',
        resolve: {
            blog: function(BlogLoader){
                return BlogLoader();
            }
        },
        templateUrl: '/app/blog/views/form.html'
    }).when('/view/:blogId', {
        controller: 'ViewCtrl',
    resolve: {
        blog: function(BlogLoader) {
            return BlogLoader();
        }
    },
    templateUrl: '/app/blog/views/view.html'
    }).when('/new', {
        controller: 'NewCtrl',
    templateUrl: '/app/blog/views/form.html'
    }).otherwise({redirectTo: '/sec'});
}]);
