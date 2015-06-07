(function(window, angular, document, undefined) {
    'use strict';
 
    angular.module('ngResourceRetries', ['ng', 'ngResource']).
    provider('resourceRetries', function(){
        var common = {
            retryTimeout : 100,
            retryMax: 3
        };
        
        this.options = function(args)
        {
          common = angular.extend(common, args);
        }
 
        // I want to override/wrap all the method generated by the $resource .. I managed don't see an another way to d
        this.$get = ['$timeout', '$q', '$resource', function($timeout, $q, $resource) {
            function resourceFactory(url, paramDefaults, actions, options) {
                // create the resource to use in our wrapper
                var ngResource = $resource(url, paramDefaults, actions);
                var res = {};
                var generateWrapper = function (methodName) {
                    return function(a1, a2, a3, a4){
                        var defer = $q.defer();
                        var inc = 0;
                        var args = arguments;
 
                        function retry(inc) { 
                            var retryMax = common.retryMax, retryTimeout = common.retryTimeout;
                            if(options && angular.isNumber(options.retryMax))
                              retryMax = options.retryMax;
                            if(options && angular.isNumber(options.retryTimeout))
                              retryTimeout = options.retryTimeout;

                            inc += 1;
                            if (inc <= retryMax) {
                                ngResource[methodName].apply(this, args).$promise.then(function(data){
                                    defer.resolve(data);                    
                                }, function(error){
                                    $timeout(function(){
                                        retry(inc);
                                    }, retryTimeout);
                                });
                            } else {
                                ngResource[methodName].apply(this, args).$promise.then(function(data){
                                    defer.resolve(data);                    
                                }, function(error){
                                    defer.reject(error);
                                });
                            }
                        }
 
                        retry(inc);
 
                        return {
                            "$promise": defer.promise
                        };
                    };
                };
 
                angular.forEach(ngResource.prototype, function(value, method){
                    // we slice(1) to delete the $
                    res[method.slice(1)] = generateWrapper(method.slice(1));
                });
 
                return res;
            }
 
            return resourceFactory;
        }]
    });
 
})(window, window.angular);