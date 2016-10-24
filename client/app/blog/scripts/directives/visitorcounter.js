/**
* @file services.js
* @brief services copied from the book
* @author yingx
* @date 2016-03-15
 */

app.directive('myTag', ['$http', function($http) {
    return {
        restrict: 'E',
        transclude: true,
        replace: true,     
        scope: {
            src:"="       
        },
        controller: function($scope) {
            console.info("enter directive controller");
            $scope.gallery = [];
            console.log($scope.src);

            $http({method: 'GET', url: $scope.src}).then(function (result) {
                console.log(result);                              
            }, function (result) {
                alert("Error: No data returned");
            });
        }
    }
}]);




.directive('ionSelect',function(){

    'use strict';
    return{
        restrict: 'EAC',
    scope: {
        label:'@',
    labelField:'@',
    provider:'=',
    ngModel: '=?',
    ngValue: '=?',
    myFunctionToCall : '='

    },
    require: '?ngModel',
    transclude : false,
    replace: false,
    template:
    '<div class="selectContainer">'
    //   +'<label class="item item-input item-stacked-label">' 
    //  +'<span class="input-label">{{label}}</span>'
    //   +'<div class="item item-input-inset">'
    +'<label class="item item-input" >'
    +'<i class="icon ion-ios-search-strong"></i>'
    +'<input id="filtro" type="search"  ng-model="ngModel" style="padding: 5px;" ng-value="ngValue" ng-keydown="onKeyDown()"/>'
    //          +'</label>'
    //     +'<button class="button button-small button-clear" ng-click="open()">'
    //          +'<i class="icon ion-chevron-down"></i>'
    //      +'</button>'
    //    +'</div>' 
    +'</label>'
    +'<div class="optionList padding-left padding-right" ng-show="showHide">'
    +'<ion-scroll>'
    +'<ul class="list">'
    +'<li class="item" ng-click="selecionar(item)" ng-repeat="item in provider | filter:ngModel" >{{item[labelField]}}</li>'                    
    +'</ul>'
    +'</ion-scroll>'
    +'</div>'    
    +'</div>'
    ,


    link: function (scope, element, attrs,ngModel) {
        scope.ngValue = scope.ngValue !== undefined ? scope.ngValue :'item';

        scope.ontouch1 = function(item){
            console.log(item[labelField]);alert(1);
        };

        scope.selecionar = function(item){
            ngModel.$setViewValue(item);
            scope.showHide = false;
        };

        element.bind('click',function(){
            element.find('input').focus();
        });

        scope.open = function(){

            scope.ngModel = "";  
            return scope.showHide=!scope.showHide;
        };

        scope.onKeyDown = function(){
            scope.showHide = true;
            if(!scope.ngModel){
                scope.showHide = false;
            }
        }

        scope.$watch('ngModel',function(newValue){
            if(newValue)

            element.find('input').val(newValue[scope.labelField]);

        if(newValue[scope.labelField]!== undefined){

            alert(newValue[scope.labelField]);

        }
        });
    },
    };  
})
