/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

app.controller('ListCtrl', ['$scope', 'blogs', function($scope, blogs){
    //console.log(blogs);
    $scope.blogs = blogs;
}]);

app.controller('ViewCtrl', ['$scope', '$location', 'blog', function($scope, $location, blog){
    $scope.blog = blog;

    $scope.edit = function(){
        $location.path('/edit/' + blog.id);
    }
}]);

app.controller('EditCtrl', ['$scope', '$location', 'blog', function($scope, $location, blog){
    $scope.blog = blog;

    $scope.save = function(){
        $scope.blog.$save(function(blog){
            $location.path('/view/' + blog.id);
        });
    };

    $scope.remove = function(){
        $scope.blog.$remove(function(blog){
            $location.path('/');
        });
    };
}]);

app.controller('NewCtrl', ['$scope', '$location', 'BlogSrv', function($scope, $location, BlogSrv){
    $scope.blog = new BlogSrv({
        id: -1
    });

    $scope.save = function(){
        $scope.blog.$save(function(blog){
            $location.path('/view/' + blog.id);
        });
    };
}]);

