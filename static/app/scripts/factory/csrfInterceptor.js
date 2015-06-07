'use strict';

angular.module('CallForPaper')
    .factory('csrfInterceptor', function() {
        /**
         * Intercep every request and add csrf token
         */
        return {
            request: function(config) {
                function b(a) {
                    return a ? (a ^ Math.random() * 16 >> a / 4).toString(16) : ([1e16] + 1e16).replace(/[01]/g, b)
                };
                // put a new random secret into our CSRF-TOKEN Cookie after each response
                document.cookie = 'CSRF-TOKEN=' + b();
                return config;
            }
        }
    })