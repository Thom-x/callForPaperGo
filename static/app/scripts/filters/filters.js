angular.module('customFilters', [])
	/**
	 * Truncate too long string to  255 chars
	 * @param  {string}
	 * @return {string}
	 */
	.filter('truncate', function() {
		return function(input) {
			if (input !== undefined) {
				if (input.length > 255)
					return input.substring(0, 255) + " ...";
				else
					return input;
			}
			return "";
		};
	});