'use strict';

angular.module('CallForPaper')
	.controller('Step2Ctrl', ['$scope', '$state', function($scope, $state) {
		$scope.$watch(function(){
			return $scope.form.sessionName.$valid && $scope.form.description.$valid && $scope.form.references.$valid && ($scope.$parent.formData.session.difficulty > 0) && $scope.form.track.$valid && $scope.form.type.$valid && $scope.form.coSpeaker.$valid;
		},function(isValid){
			$scope.$parent.formData.steps.isValid[1]= isValid;
		})

		$scope.verify = false;
		$scope.doVerify = function()
		{
			$scope.verify = true;
			if($scope.$parent.formData.steps.isValid[1] &&  $scope.$parent.formData.steps.isValid[0])
			{
				$state.go('app.form.step3');
			}
		}
	}]);
