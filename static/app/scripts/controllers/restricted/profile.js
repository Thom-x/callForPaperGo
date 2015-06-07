'use strict';

angular.module('CallForPaper')
	.controller('ProfileCtrl', ['$scope', 'RestrictedUser', function($scope, RestrictedUser) {
		$scope.formData = {};
		$scope.formData.phone = "";
		$scope.$watch(function() {
			return $scope.form.name.$valid && $scope.form.firstname.$valid && ($scope.form.phone.$valid || $scope.formData.phone == "") && $scope.form.company.$valid && $scope.form.bio.$valid && $scope.form.social.$valid;
		}, function(isValid) {
			$scope.formData.isValid = isValid;
		})

		$scope.$watch(function() {
			if ($scope.formData.socialArray !== undefined)
				return $scope.formData.socialArray.length;
		}, function() {
			if ($scope.formData.socialArray !== undefined) {
				$scope.formData.social = $scope.formData.socialArray.map(function(elem) {
					return elem.text;
				}).join(", ");
			}
		})

		RestrictedUser.get(function(profile) {
			if (profile !== undefined) {
				for (var key in profile) {
					if (profile.hasOwnProperty(key)) {
						switch (key) {
							case "bio":
								if (profile[key] !== null) $scope.formData.bio = profile[key];
								break;
							case "company":
								if (profile[key] !== null) $scope.formData.company = profile[key];
								break;
							case "firstname":
								if (profile[key] !== null) $scope.formData.firstname = profile[key];
								break;
							case "name":
								if (profile[key] !== null) $scope.formData.name = profile[key];
								break;
							case "phone":
								if (profile[key] !== null) $scope.formData.phone = profile[key];
								break;
							case "social":
								if (profile[key] !== null && profile[key] != "") $scope.formData.socialArray = profile[key].split(", ").map(function(value) {
									return {
										text: value
									};
								});
								break;
						}
					}
				}
			}
		})

		$scope.verify = false;
		$scope.sendError = false;
		$scope.sendSuccess = false;
		$scope.doVerify = function() {
			$scope.verify = true;
			if ($scope.formData.isValid) {
				// put
				$scope.sending = true;
				RestrictedUser.update({}, $scope.formData, function(success) {
					$scope.sendSuccess = true;
					$scope.sendError = false;
					$scope.sending = false;
				}, function(error) {
					$scope.sendSuccess = false;
					$scope.sendError = true;
					$scope.sending = false;
				});
			}
		}
	}]);