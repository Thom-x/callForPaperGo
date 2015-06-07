angular.module('CallForPaper').factory('Application', ['$resource', function($resource) {
	return $resource('api/application/:id', null, {
	});
}]);