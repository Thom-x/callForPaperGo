'use strict';

angular.module('CallForPaper')
	.controller('AdminSessionsCtrl', ['$scope', 'AdminSession', '$filter', 'ngTableParams', '$q', 'Notification', 'screenSize', 'AdminUser', 'Config', function($scope, AdminSession, $filter, ngTableParams, $q, Notification, screenSize, AdminUser, Config) {
		var sessions = []
		$scope.gcmApiKey = Config.gcmApiKey;
		$scope.screenSize = screenSize;
		$scope.realDifficulty = [$filter('translate')('step2.beginner'), $filter('translate')('step2.confirmed'), $filter('translate')('step2.expert')];
		AdminSession.query().$promise.then(function(sessionsTmp) {
			sessions = sessionsTmp.map(function(session) {
				session.fullname = session.name + " " + session.firstname;
				session.keyDifficulty = (['beginner', 'confirmed', 'expert'])[session.difficulty - 1];
				return session;
			});
			updateTable();
		});

		$scope.difficulties = function(column) {
			var def = $q.defer();
			var difficulties = [{
				'id': '1',
				'title': $filter('translate')('step2.beginner')
			}, {
				'id': '2',
				'title': $filter('translate')('step2.confirmed')
			}, {
				'id': '3',
				'title': $filter('translate')('step2.expert')
			}];
			def.resolve(difficulties);
			return def;
		};

		$scope.tracks = function(column) {
			var def = $q.defer();
			var difficulties = [{
				'id': 'web',
				'title': $filter('translate')('step2.tracks.web')
			}, {
				'id': 'discovery',
				'title': $filter('translate')('step2.tracks.discovery')
			}, {
				'id': 'mobile',
				'title': $filter('translate')('step2.tracks.mobile')
			}, {
				'id': 'cloud',
				'title': $filter('translate')('step2.tracks.cloud')
			}];
			def.resolve(difficulties);
			return def;
		};

		$scope.gcmSend = function(sub)
		{
			AdminUser.postNotifToken({"token" : sub.subscriptionId})
		}

		var updateTable = function() {
			$scope.tableParams = new ngTableParams({
				count: 10,
				filter: {
					nafullnameme: '', // initial filter
					description: '',
					difficulty: '',
					track: ''
				},
				sorting: {
					mean: 'desc'
				}
			}, {
				filterDelay : 0,
				total: sessions.length, // length of data
				getData: function($defer, params) {
					// use build-in angular filter
					var orderedData = params.filter() ?
						$filter('filter')(sessions, params.filter()) : sessions;
					orderedData = params.sorting() ? $filter('orderBy')(orderedData, params.orderBy()) : orderedData;

					$scope.sessions = orderedData.slice((params.page() - 1) * params.count(), params.page() * params.count());

					params.total(orderedData.length); // set total for recalc pagination
					$defer.resolve($scope.sessions);
				}
			});
		}
	}]);