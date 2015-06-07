angular.module('CallForPaper').factory('AdminComment', ['$resource', function($resource) {
  return $resource('api/admin/comment/:id',{},
  	{
  		getAll :{
  			method: 'GET',
  			isArray: true
  		},
  		getByRowId :{
  			url: 'api/admin/comment/row/:rowId',
  			method: 'GET',
  			isArray: true
  		},
  		update: { method:'PUT' }
  	});
}]);