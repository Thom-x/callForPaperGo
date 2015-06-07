'use strict';

angular.module('CallForPaper')
	.controller('RestrictedSessionCtrl', ['$scope', '$stateParams', '$filter', 'RestrictedSession', function($scope, $stateParams, $filter, RestrictedSession) {
		$scope.session = null;
		RestrictedSession.get({
			id: $stateParams.id
		}).$promise.then(function(sessionTmp) {
			$scope.session = sessionTmp;
			$scope.session.keyDifficulty = (['beginner', 'confirmed', 'expert'])[sessionTmp.difficulty - 1];
		});
	}]);