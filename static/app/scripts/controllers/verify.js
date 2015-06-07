'use strict';

angular.module('CallForPaper')
	.controller('VerifyCtrl', ['$scope', '$auth', 'Verify', '$stateParams', function($scope, $auth, Verify, $stateParams) {
        $scope.verified = false;
        $scope.alreadyVerified = false;
        $scope.notVerified = false;

		Verify.get($stateParams.id, $stateParams.token)
		.success(function(data, status) {
            $auth.setToken(data.token, true);
            $scope.verified = true;
        })
        .error(function(data, status) {
        	if(status === 409)
        	{
        		$scope.alreadyVerified = true;
        	} else {
        		$scope.notVerified = true;
        	}
        });
	}]);