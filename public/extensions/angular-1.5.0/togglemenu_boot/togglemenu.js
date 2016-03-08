angular.module('util.treeMenu', [])
.directive("menuLink",function() {
    return {
        scope: {
        	menu: "="
        },
        template: '\
        	<a class="md-button" ng-class="{\'active\' : isSelected()}" ng-click="focusSection()">\
        		<i class="fa {{menu.icon}} fw"></i> {{menu.name}}\
		    </a>\
        ',
        link: function($scope, $element) {
			var controller = $element.parent().controller();
			$scope.isSelected = function() {
				return controller.isSelected($scope.menu);
			};
			$scope.focusSection = function() {
				controller.menuClick($scope.menu);
			};
	    }
    }
})
.directive('menuToggle', [ '$timeout', function($timeout) {
    return {
        scope: {
        	menu: "="
        },
        template: '\
        	<a class="md-button md-button-toggle" ng-class="{\'active\' : isSelected()}" ng-click="toggle()">\
        		<i class="fa {{menu.icon}} fw"></i>  {{menu.name}} <i class="pull-right md-toggle-icon fa fa-angle-down" ng-class="{\'toggled\' : isOpen()}"></i>\
		    </a>\
	        <ul class="menu-toggle-list">\
	        	<li ng-repeat="subMenu in menu.subMenus">\
	        		<menu-link menu="subMenu"></menu-link>\
	        	</li>\
	        </ul>\
        ',
        link: function($scope, $element) {
	      var controller = $element.parent().controller();
	
	      $scope.isOpen = function() {
	        return controller.isOpen($scope.menu);
	      };
	      $scope.toggle = function() {
	        controller.toggleOpen($scope.menu);
	      };
	      $scope.$watch(
	          function () {
	            return controller.isOpen($scope.menu);
	          },
	          function (open) {
	            var $ul = $element.find('ul');
	            var targetHeight = open ? getTargetHeight() : 0;
	            $timeout(function () {
	              $ul.css({ height: targetHeight + 'px' });
	            }, 0, false);
	
	            function getTargetHeight () {
	              var targetHeight;
	              $ul.addClass('no-transition');
	              $ul.css('height', '');
	              targetHeight = $ul.prop('clientHeight');
	              $ul.css('height', 0);
	              $ul.removeClass('no-transition');
	              return targetHeight;
	            }
	          }
	      );
	      var parentNode = $element[0].parentNode.parentNode.parentNode;
	      if(parentNode.classList.contains('parent-list-item')) {
	        var heading = parentNode.querySelector('h2');
	        $element[0].firstChild.setAttribute('aria-describedby', heading.id);
	      }
	    }
    }
}]);
