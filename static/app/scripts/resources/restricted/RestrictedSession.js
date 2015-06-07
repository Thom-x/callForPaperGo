angular.module('CallForPaper').factory('RestrictedSession', ['resourceRetries', function(resourceRetries) {
  return resourceRetries('api/restricted/session/:id', null, 
  	{
        update: { method:'PUT' , url: 'api/restricted/session/:id' },
    });
}]);