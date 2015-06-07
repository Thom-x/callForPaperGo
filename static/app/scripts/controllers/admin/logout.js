'use strict';

angular.module('CallForPaper')
	.controller('AdminLogoutCtrl', ['$scope', '$interval', 'AuthService', function($scope, $interval, AuthService) {
		$scope.dots = "...";
		var updateDots = function() {
			$scope.dots = $scope.dots + ".";
			if ($scope.dots.length === 4) {
				$scope.dots = "";
				return;
			}
		}
		var loading = $interval(updateDots, 1000);
		AuthService.logout('app.login');
	}]);