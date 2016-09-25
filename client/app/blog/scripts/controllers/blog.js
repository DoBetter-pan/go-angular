/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

blog.controller('IndexCtrl', ['$scope', 'articles', function($scope, articles){
    //console.log(articles);
    $scope.articles = articles;
}]);

blog.controller('ListCtrl', ['$scope', 'articles', function($scope, articles){
    //console.log(articles);
    $scope.articles = articles;
}]);

blog.controller('ViewCtrl', ['$scope', '$location', 'article', function($scope, $location, article){
    $scope.article = article;

    $scope.edit = function(){
        $location.path('/edit/' + article.id);
    }
}]);

blog.controller('EditCtrl', ['$scope', '$location', 'article', function($scope, $location, article){
    $scope.article = article;

    $scope.save = function(){
        $scope.article.$save(function(article){
            $location.path('/view/' + article.id);
        });
    };

    $scope.remove = function(){
        $scope.article.$remove(function(article){
            $location.path('/');
        });
    };
}]);

blog.controller('NewCtrl', ['$scope', '$location', 'BlogSrv', function($scope, $location, BlogSrv){
    $scope.article = new BlogSrv({
        id: -1
    });

    $scope.save = function(){
        $scope.article.$save(function(article){
            $location.path('/view/' + article.id);
        });
    };
}]);

