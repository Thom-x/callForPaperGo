'use strict';

angular.module('CallForPaper')
	.provider('$appConfig', function() {
		this.$get = function()
		{
			var $appConfig = {};
			$appConfig.isConfigured = ['Application', '$q', '$state', function(Application, $q, $state) {
				var deferred = $q.defer();
				Application.get(function(config) {
					if (config.configured === true) {
						deferred.reject();
						$state.go("config");
					} else {
						deferred.resolve();
					}
				})
				return deferred.promise;
			}];
			return $appConfig;
		}
	});