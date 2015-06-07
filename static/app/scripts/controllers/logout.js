'use strict';

angular.module('CallForPaper')
	.controller('LogoutCtrl', ['$auth', '$state', '$scope', function($auth, $state, $scope) {
		if (!$auth.isAuthenticated()) {
			$scope.$emit('authenticate');
			$state.go('app.login');
			return;
		}
		$auth.logout().then(function(){
			$scope.$emit('authenticate');
			$state.go('app.login');
		});
	}]);