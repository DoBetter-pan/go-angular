var mainApp = angular.module("mainApp", ['ngAside']);
mainApp.config(function($interpolateProvider){
	$interpolateProvider.startSymbol('[[').endSymbol(']]');
});
mainApp.controller("mainController", function($scope){
	$scope.navLock = false;
	$scope.isRouteLoading = false;
	$scope.openAside = function(){
		console.log("openAside...");
		$scope.isRouteLoading = !$scope.isRouteLoading;
	}
});
