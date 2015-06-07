'use strict';

angular.module('CallForPaper')
	.controller('HeaderCtrl', ['$scope', '$rootScope', '$translate', '$auth', 'Application', function($scope, $rootScope, $translate, $auth, Application) {
		$scope.language = $translate.use();

		$scope.changeLanguage = function(key) {
			$translate.use(key);
		};

		$rootScope.$on('$translateChangeEnd', function(event, args) {
			$scope.language = args.language;
		});

		$scope.isAuthenticated = $auth.isAuthenticated();
		$scope.$on('authenticate', function() { $scope.isAuthenticated = $auth.isAuthenticated(); });

		/**
		 * Get eventName
		 */
		Application.get(function(config) {
			$scope.title = config.eventName;
		})
	}]);