angular.module('CallForPaper').factory('Verify', ['$http', function($http) {
	return {
		get: function(id, token) {
			return $http({
				url: '/auth/verify',
				method: "GET",
				params: {
					id: id,
					token: token
				}
			});
		}
	}
}]);