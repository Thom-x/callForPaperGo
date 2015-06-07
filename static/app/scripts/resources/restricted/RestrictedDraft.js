angular.module('CallForPaper').factory('RestrictedDraft', ['resourceRetries', function(resourceRetries) {
	return resourceRetries('api/restricted/draft/:id', null, {
        update: { method:'PUT' , url: 'api/restricted/draft/:id' },
		delete: {
			method: 'DELETE',
			url: 'api/restricted/draft/:id'
		}
	});
}]);