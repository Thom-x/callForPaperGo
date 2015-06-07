angular.module('CallForPaper').factory('AdminRate',['$resource', function($resource) {
  return $resource('api/admin/rate/:id', null, 
  	{
        update: { method:'PUT', url: 'api/admin/rate/:id' },
        getByRowIdAndUserId: {
        	method:'GET',
        	url: 'api/admin/rate/row/:rowId/user/me'
        },
        getByRowId: {
        	method:'GET',
        	url: 'api/admin/rate/row/:rowId',
        	isArray: true
        }
    });
}]);