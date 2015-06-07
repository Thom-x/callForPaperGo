'use strict';

angular.module('CallForPaper')
    .factory('authHttpResponseInterceptor', ['$q', '$injector', '$filter', function($q, $injector, $filter) {
        /**
         * Intercep every request and popup a notification if error
         */
        return {
            response: function(response) {
                return response || $q.when(response);
            },
            responseError: function(rejection) {
                var Notification = $injector.get('Notification');
                if (rejection.status === 0) {
                    Notification.error({
                        message: $filter('translate')('error.noInternet'),
                        delay: 3000
                    });

                } else if (rejection.status === 401) {
                    $injector.get('$state').go('app.login');
                } else if (rejection.status === 403) {
                    $injector.get('$state').go('403');
                } else if (rejection.status === 404) {
                    $injector.get('$state').go('404');
                } else if (rejection.status === 409) {
                    // Nothing
                } else if(rejection.status === 500 && rejection.data.error === "Internal Server Error") {
                    // Nothing
                } else {
                    Notification.error({
                        message: $filter('translate')('error.backendcommunication'),
                        delay: 3000
                    });
                }
                return $q.reject(rejection);
            }
        }
    }])