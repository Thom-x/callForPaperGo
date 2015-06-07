'use strict';

angular.module('CallForPaper')
	.controller('AdminSessionCtrl', ['$scope', '$stateParams', '$filter', '$translate', 'AdminSession', 'AdminComment', 'AdminRate', function($scope, $stateParams, $filter, $translate, AdminSession, AdminComment, AdminRate) {
		$scope.session = null;
		AdminSession.get({
			id: $stateParams.id
		}).$promise.then(function(sessionTmp) {
			$scope.session = sessionTmp;
			$scope.session.keyDifficulty = (['beginner', 'confirmed', 'expert'])[sessionTmp.difficulty - 1];
		});
		/**
		 * get comments of the session
		 * @return {[AdminComment]}
		 */
		var updateComments = function() {
			AdminComment.getByRowId({
				rowId: $stateParams.id
			}, function(commentsTmp) {
				$scope.comments = commentsTmp;
			})
		}
		updateComments();

		$scope.commentButtonDisabled = false;
		$scope.postComment = function() {
			$scope.commentButtonDisabled = true;
			AdminComment.save({
				'comment': $scope.commentMsg,
				'rowId': $stateParams.id
			}, function(c) {
				$scope.commentMsg = "";
				$scope.commentButtonDisabled = false;
				updateComments();
			});
		}

		/**
		 * get rates of the session
		 * @return {[AdminRate]}
		 */
		var updateRates = function() {
			AdminRate.getByRowId({
				'rowId': $stateParams.id
			}, function(ratesTmp) {
				$scope.rates = ratesTmp;
				$scope.mean = $scope.rates.map(function(rateTmp) {
					return rateTmp.rate;
				}).reduce(function(x, y) {
					return x + y;
				}, 0) / ($scope.rates.length == 0 ? 1 : $scope.rates.length);
			})
		}
		updateRates();

		$scope.yourRate = {
			rate: 0,
			id: undefined
		};
		/**
		 * Obtain current user rate
		 * @param  {long : rowId}
		 * @return {AdminRate}
		 */
		AdminRate.getByRowIdAndUserId({
			'rowId': $stateParams.id,
		}, function(rateTmp) {
			if (rateTmp.id !== undefined) {
				$scope.yourRate = rateTmp;
			}
		})

		$scope.rateButtonDisabled = false;
		/*
		 *	Post new rate
		 */
		$scope.postRate = function() {
			$scope.rateButtonDisabled = true;
			if ($scope.yourRate.id === undefined) {
				AdminRate.save({
					'rate': $scope.yourRate.rate,
					'rowId': $stateParams.id
				}, function(c) {
					$scope.yourRate.id = c.id;
					$scope.rateButtonDisabled = false;
					updateRates();
				});
			} else {
				AdminRate.update({
					'id': $scope.yourRate.id
				}, {
					'rate': $scope.yourRate.rate,
					'rowId': $stateParams.id
				}, function(c) {
					$scope.rateButtonDisabled = false;
					updateRates();
				});
			}
		}
	}]);