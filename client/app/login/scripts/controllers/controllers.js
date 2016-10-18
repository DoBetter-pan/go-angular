/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

app.controller('ViewCtrl', ['$scope', '$location', 'login', function($scope, $location, login){
    $scope.login = login;

    $scope.check = function(){
        $location.path('/edit/' + login.id);
    }
}]);

app.controller('NewCtrl', ['$scope', '$location', 'LoginSrv', function($scope, $location, LoginSrv){
    $scope.login = new LoginSrv({
        id: -1
    });

    $scope.save = function(){
        $scope.login.$save(function(login){
            $location.path('/view/' + login.id);
        });
    };
}]);

