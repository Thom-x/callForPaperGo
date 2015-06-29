'use strict';

angular.module('CallForPaper')
	.controller('FormCtrl', ['$scope', '$filter', '$translate', 'RestrictedSession', 'RestrictedDraft', 'RestrictedUser', '$state', '$stateParams', function($scope, $filter, $translate, RestrictedSession, RestrictedDraft, RestrictedUser, $state, $stateParams) {
		// we will store all of our form data in this object
		$scope.formData = {};
		$scope.formData.steps = {};
		$scope.formData.steps.currentStep = 1;
		$scope.formData.sending = false;
		$scope.formData.steps.isValid = [false, false, false];
		$scope.language = $translate.use();

		$scope.formData.speaker = {};
		$scope.formData.speaker.phone = "";
		$scope.formData.session = {};
		$scope.formData.help = {};

		var id = $stateParams.id;
		var parseRow = function(id)
		{
			if (id !== "") {
				RestrictedDraft.get({
					id: id
				}).$promise.then(function(draft) {
					if(draft.added !== undefined)
					{
						for (var key in draft) {
							if (draft.hasOwnProperty(key)) {
								switch (key) {
									case "bio":
										if (draft[key] !== null) $scope.formData.speaker.bio = draft[key];
										break;
									case "coSpeaker":
										if (draft[key] !== null) $scope.formData.session.coSpeaker = draft[key];
										break;
									case "company":
										if (draft[key] !== null) $scope.formData.speaker.company = draft[key];
										break;
									case "description":
										if (draft[key] !== null) $scope.formData.session.description = draft[key];
										break;
									case "difficulty":
										if (draft[key] !== null) { $scope.formData.session.difficulty = draft[key]; $scope.hoverDifficulty(draft[key]); }
										break;
									case "email":
										if (draft[key] !== null) $scope.formData.speaker.email = draft[key];
										break;
									case "financial":
										if (draft[key] !== null) $scope.formData.help.financial = draft[key];
										break;
									case "firstname":
										if (draft[key] !== null) $scope.formData.speaker.firstname = draft[key];
										break;
									case "hotel":
										if (draft[key] !== null) $scope.formData.help.hotel = draft[key];
										break;
									case "hotelDate":
										if (draft[key] !== null) $scope.formData.help.hotelDate = draft[key];
										break;
									case "name":
										if (draft[key] !== null) $scope.formData.speaker.name = draft[key];
										break;
									case "phone":
										if (draft[key] !== null) $scope.formData.speaker.phone = draft[key];
										break;
									case "references":
										if (draft[key] !== null) $scope.formData.session.references = draft[key];
										break;
									case "sessionName":
										if (draft[key] !== null) $scope.formData.session.sessionName = draft[key];
										break;
									case "social":
										if (draft[key] !== null && draft[key] != "") $scope.formData.speaker.socialArray = draft[key].split(", ").map(function(value){ return {text : value}; });
										break;
									case "track":
										if (draft[key] !== null) $scope.formData.session.track = draft[key];
										break;
									case "type":
										if (draft[key] !== null) $scope.formData.session.type = draft[key];
										break;
									case "travel":
										if (draft[key] !== null) $scope.formData.help.travel = draft[key];
										break;
									case "travelFrom":
										if (draft[key] !== null) $scope.formData.help.travelFrom = draft[key];
										break;
								}
							}
						}
					}
					else
					{
						// not existing
						$state.go("404");
					}
				})
			}
			else
			{
				RestrictedUser.query(function(profile) {
					if (profile !== undefined) {
						for (var key in profile) {
							if (profile.hasOwnProperty(key)) {
								switch (key) {
									case "bio":
										if (profile[key] !== null) $scope.formData.speaker.bio = profile[key];
										break;
									case "company":
										if (profile[key] !== null) $scope.formData.speaker.company = profile[key];
										break;
									case "email":
										if (profile[key] !== null) $scope.formData.speaker.email = profile[key];
										break;
									case "firstname":
										if (profile[key] !== null) $scope.formData.speaker.firstname = profile[key];
										break;
									case "name":
										if (profile[key] !== null) $scope.formData.speaker.name = profile[key];
										break;
									case "phone":
										if (profile[key] !== null) $scope.formData.speaker.phone = profile[key];
										break;
									case "social":
										if (profile[key] !== null && profile[key] != "") $scope.formData.speaker.socialArray = profile[key].split(", ").map(function(value) {
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
			}
		}
		parseRow(id);
		
		/**
		 * Send the form to the server
		 * @param  {Boolean}
		 * @return {void}
		 */
		$scope.processForm = function(isValid) {
			$scope.formData.sending = true;
			var model = {};
			angular.extend(model, $scope.formData.help);
			angular.extend(model, $scope.formData.speaker);
			angular.extend(model, $scope.formData.session);

			if (id !== "") {
				// put
				RestrictedSession.update({id : id}, model).$promise.then(function(success) {
					$scope.formData.sending = false;
					$state.go('app.form.result');
				}, function(error) {
					$scope.formData.sending = false;
					$scope.sendError = true;
				});
			}
			else
			{
				// save
				RestrictedSession.save(model).$promise.then(function(success) {
					$scope.formData.sending = false;
					$state.go('app.form.result');
				}, function(error) {
					$scope.formData.sending = false;
					$scope.sendError = true;
				});
			}
		};

		/**
		 * Send the draft to the server
		 * @param  {Boolean}
		 * @return {void}
		 */
		$scope.processSaveForm = function() {
			$scope.formData.sending = true;
			var model = {};
			angular.extend(model, $scope.formData.help);
			angular.extend(model, $scope.formData.speaker);
			angular.extend(model, $scope.formData.session);

			// empty previous completed field
			for (var key in model) {
				if(model[key] === undefined)
				{
					model[key] = "";
				}
			}

			if (id !== "") {
				// put
				RestrictedDraft.update({id : id}, model).$promise.then(function(success) {
					$scope.formData.sending = false;
					$state.go('app.dashboard');
				}, function(error) {
					$scope.formData.sending = false;
					$scope.sendError = true;
				});
			}
			else
			{
				// save
				RestrictedDraft.save(model).$promise.then(function(success) {
					$scope.formData.sending = false;
					$state.go('app.dashboard');
				}, function(error) {
					$scope.formData.sending = false;
					$scope.sendError = true;
				});
			}
		};

		/**
		 * Update difficulty label when over
		 */
		$scope.hoverDifficulty = function(value) {
			$scope.formData.session.difficultyLabel = ([$filter('translate')('step2.beginner'), $filter('translate')('step2.confirmed'), $filter('translate')('step2.expert')])[value - 1];
		};
	}]);