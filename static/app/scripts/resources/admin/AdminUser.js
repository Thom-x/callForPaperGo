angular.module('CallForPaper').factory('AdminUser', ['$resource', function($resource) {
	return $resource('/api/commonAdmin/:id', null, {
		getCurrentUser: {
			url: '/api/commonAdmin/currentUser',
			method: 'GET'
		},
		getLoginUrl: {
			url: '/api/commonAdmin/login',
			method: 'POST'
		},
		getLogoutUrl: {
			url: '/api/commonAdmin/logout',
			method: 'POST'
		},

		postNotifToken: {
			url: '/api/admin/user/notif',
			method: 'POST'
		}
	});
}]);