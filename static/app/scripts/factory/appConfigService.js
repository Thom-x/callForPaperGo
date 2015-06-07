'use strict';

angular.module('CallForPaper')
  .factory('AppConfig', function() {
    var appConfigService = {};
    /**
     * Redirect user if app not configured
     */
    appConfigService.isConfigured = ['Application', '$q', '$state', function(Application, $q, $state) {
      var deferred = $q.defer();
      Application.get(function(config) {
        if (config.configured === false) {
          deferred.reject();
          $state.go("config.about");
        } else {
          deferred.resolve();
        }
      })
      return deferred.promise;
    }]
    return appConfigService;
  })