angular.module('CallForPaper').factory('AdminSession', ['resourceRetries', function(resourceRetries) {
  return resourceRetries('api/admin/session/:id', null, 
  	{
        update: { method:'PUT' },
    });
}]);