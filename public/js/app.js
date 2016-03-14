var mainApp = angular.module("mainApp", ['ngAside']);
mainApp.config(function($interpolateProvider){
	$interpolateProvider.startSymbol('[[').endSymbol(']]');
});
mainApp.controller("mainController", function($scope){
	var menuList = eval('(' + '[{"icon":"fa-folder","subMenus":[{"icon":"fa-file-o","genFile":{"nodeId":"21","projectId":"2","type":"ngjs"},"name":"用户管理","path":"admin/User","type":"link"},{"icon":"fa-file-o","genFile":{"nodeId":"22","projectId":"2","type":"ngjs"},"name":"客户管理","path":"admin/Customer","type":"link"}],"name":"后台管理","type":"toggle"},{"icon":"fa-file-o","genFile":{"nodeId":"23","projectId":"2","type":"ngjs"},"name":"员工客户统计","path":"module/UserCustomer","type":"link"}]' + ')');
	
	console.log(menuList);
	$scope.navLock = true;
	$scope.isRouteLoading = true;
	$scope.menu = {
		menus: menuList
	};
	$scope.xxx = {
		xxxx: ["C", "Java"]
	};
	$scope.openAside = function(){
		console.log("openAside...");
		$scope.isRouteLoading = !$scope.isRouteLoading;
	}
});
