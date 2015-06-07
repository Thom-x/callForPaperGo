'use strict';

angular.module('CallForPaper')
	.controller('DashboardCtrl', ['$scope', '$filter', 'RestrictedSession', 'RestrictedDraft', 'AuthService', function($scope, $filter, RestrictedSession, RestrictedDraft, AuthService) {
		$scope.sessions = [];
		$scope.sessionsLoaded = false;
		$scope.realDifficulty = [$filter('translate')('step2.beginner'), $filter('translate')('step2.confirmed'), $filter('translate')('step2.expert')];
		var querySession = function() {
			RestrictedSession.query(function(sessionsTmp) {
				$scope.sessions = sessionsTmp.map(function(session) {
					session.fullname = session.name + " " + session.firstname;
					session.keyDifficulty = (['beginner', 'confirmed', 'expert'])[session.difficulty - 1];
					return session;
				});
				$scope.sessionsLoaded = true;
			});
		}

		$scope.drafts = [];
		$scope.draftsLoaded = false;
		var queryDraft = function() {
			RestrictedDraft.query().$promise.then(function(draftsTmp) {
				$scope.drafts = draftsTmp
				$scope.draftsLoaded = true;
			});
		}

		$scope.delete = function(added) {
			RestrictedDraft.delete({
				id: added
			}).$promise.then(function() {
				queryDraft();
			});
		}

		$scope.isVerified = AuthService.isVerified();
		if ($scope.isVerified) {
			queryDraft();
			querySession();
		}

		$scope.konamiCode = false;
		$scope.launchKonami = function()
		{
			$scope.konamiCode = !$scope.konamiCode;
			$scope.$apply();
		}
	}]);