angular.module('bs-has', [])
  .factory('bsProcessValidator', ['$timeout', function($timeout) {
    return function(scope, element, ngClass, bsClass) {
      $timeout(function() {
        var input = element.find('input');
        if(!input.length) { input = element.find('select'); }
        if(!input.length) { input = element.find('textarea'); }
        if(!input.length) { input = element.find('span[rating]'); }
        if (input.length) {
            scope.$watch(function() {
                return input.hasClass(ngClass) && (input.hasClass('ng-dirty') || input.hasClass('ng-verify'));
            }, function(isValid) {
                element.toggleClass(bsClass, isValid);
            });
        }
      });
    };
  }])
  .factory('bsSubmitValidator', ['$timeout', function($timeout) {
    return function(scope, element, ngClass, bsClass, state) {
      $timeout(function() {
        var input = element.find('input');
        if(!input.length) { input = element.find('select'); }
        if(!input.length) { input = element.find('textarea'); }
        if(!input.length) { input = element.find('span[rating]'); }
        if (input.length) {
          scope.$watch(element.attr('verify'),function(value){
              if(value)
              {
                element.toggleClass(bsClass, input.hasClass(ngClass));
              }
          })
        }
      },100);
    };
  }])
  .directive('bsHasSuccess', ['bsProcessValidator', function(bsProcessValidator) {
    return {
      restrict: 'A',
      link: function(scope, element) {
        bsProcessValidator(scope, element, 'ng-valid', 'has-success');
      }
    };
  }])
  .directive('bsHasError', ['bsProcessValidator', function(bsProcessValidator) {
    return {
      restrict: 'A',
      link: function(scope, element) {
        bsProcessValidator(scope, element, 'ng-invalid', 'has-error');
      }
    };
  }])
  .directive('bsHas', ['bsProcessValidator', 'bsSubmitValidator', function(bsProcessValidator, bsSubmitValidator) {
    return {
      restrict: 'A',
      link: function(scope, element, attrs) {
        bsProcessValidator(scope, element, 'ng-valid', 'has-success');
        bsProcessValidator(scope, element, 'ng-invalid', 'has-error');
        bsSubmitValidator(scope, element, 'ng-invalid', 'has-error');
      }
    };
  }]);