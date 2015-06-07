'use strict';

angular.module('CallForPaper')
  .config(['$translateProvider', function($translateProvider) {
    $translateProvider.translations('fr-FR', {
      "lang": {
        "fr-FR": "Français",
        "en-US": "Anglais"
      },
      "header": {
        "login": "Se connecter",
        "logout": "Se déconnecter",
        "dashboard": "Dashboard",
        "profile" : "Profile"
      },
      "dashboard" :
      {
        "createNewSession" : "Créer un nouveau talk",
        "editionTalks" : "Talks en cours d'édition",
        "noEditionTalks" : "Vous n'avez pas de talk en cours d'édition",
        "sendedTalks" : "Talks envoyés",
        "draftModifiedAt" : "Brouillon modifié le ",
        "postedAt" : "ajouté le ",
        "noSendedTalks" : "Vous n'avez pas encore proposé de talk",
        "verification" : "Vérification",
        "verificationNeeded" : "Un e-mail a été envoyé à votre adresse, cliquez sur le lien présent dans celui-ci pour valider votre compte.",
      },
      "profile" :
      {
        "save" : "Sauvegarder",
        "success" : "Profile sauvegardé",
        "error" : "Erreur lors de l'enregistrement",
      },
      "login" : {
        "email" : "Email",
        "password" : "Mot de passe",
        "login" : "Se connecter",
        "noAccount" : "Vous n'avez pas encore de compte ?",
        "signup" : "S'enregistrer",
        "or" : "Ou",
        "signGoogle" : "Se connecter avec Google",
        "signGithub" : "Se connecter avec Github",
        "wait": "Veuillez patienter",
        "badCredentials": "Login ou mot de passe incorrect",
        "alreadyLinked" : "Il existe déjà un utilisateur associé à ce fournisseur"
      },
      "signup" : {
        "signup" : "S'enregistrer",
        "email" : "Email",
        "password" : "Mot de passe",
        "confirmPassword" : "Confirmez le mot de passe",
        "yesAccount" : "Vous avez déjà un compte ?",
        "loginNow" : "Se connecter",
        "emailRequired" : "L'adresse email est obligatoire",
        "emailPattern" : "Votre adresse email est invalide",
        "passwordRequired" : "Le mot de passe est obligatoire",
        "passwordMinLength" : "Le mot de passe doit contenir au moins 6 caractères",
        "passwordMatch" : "Les mots de passe doivent être identiques",
        "alreadyExists" : "Il existe déjà un utilisateur associé à cet adresse e-mail"
      },
      "verify" : {
        "title" : "Verification de l'email",
        "verified" : "Votre email a été vérifié",
        "alreadyVerified" : "Ce compte a déjà été vérifié",
        "notVerified" : "Erreur de vérification"
      },
      "step1": {
        "email": "Adresse email * :",
        "name": "Nom * :",
        "firstname": "Prénom * :",
        "phone": "Téléphone :",
        "errorPhone": "Entrez un numéro de téléphone correct",
        "company": "Entreprise :",
        "bio": "Bio * :",
        "hintBio": "Décrivez vous en quelques mots. Cette description sera utilisée sur le site web.",
        "social": "URL (Google, Github, etc) :",
        "hintSocial": "Donnez les liens de vos réseaux sociaux (pour le site web) : Twitter / G+ / Github / Blog."
      },
      "step2": {
        "name": "Nom de la conférence * :",
        "description": "Description * :",
        "hintDescription": "Donnez une description de votre présentation. Elle sera utilisée sur le site web.",
        "references": "Références ou compléments d'informations :",
        "hintReferences": "Y a-t-il des conférences où vous avez déjà fait des présentations ? Si vous pouvez donner un lien vers celle(s)-ci ça serait bien.",
        "difficulty": "Difficulté * (Débutant, Confirmé, Expert) :",
        "track": "Track * :",
        "cospeaker": "Co-conférenciers :",
        "hintCospeaker": "Si vous n'êtes pas seul lors de la présentation, donnez les nom / email / bio / liens sociaux des autres conférenciers.",
        "beginner": "Débutant",
        "confirmed": "Confirmé",
        "expert": "Expert",
        "tracks": {
          "cloud": "Cloud",
          "mobile": "Objets connectés",
          "web": "Web",
          "discovery": "Découverte"
        },
        "type": "Type * :",
        "hintType": "Une conférence doit durer environ 45 minutes et un codelab 2 heures.",
        "types": {
          "conference": "Conférence",
          "codelab": "Codelab"
        },
        "hintTrack": "Choisissez la catégorie dans laquelle vous pensez que votre conférence se situe."
      },
      "step3": {
        "header1": "Vous pouvez renseigner ici les informations nécessaires pour votre venue. Ces informations seront minutieusement étudiées pour notre décision. Ne choisissez oui que si vous en avez besoin.",
        "header2": "Un petit déjeuner et un déjeuner sont offerts le jour de l'évènement.",
        "financial": "Avez-vous besoin d’une aide financière ? * :",
        "labelTravel": "Voyage :",
        "travel": "J'ai besoin d’une aide financière pour le voyage.",
        "place": "D’où venez vous ? :",
        "labelHotel": "Hébergement",
        "hotel": "J'ai besoin d’une aide financière pour l’hotel.",
        "date": "Pour quelle(s) date(s) ? :",
        "sendError": "Erreur lors de l'envoi veuillez réessayer"
      },
      "confirmModal": {
        "title" : "Confirmation",
        "text" : "Êtes-vous sûr de vouloir envoyer ce talk ? Une fois envoyé, vous ne serez pas en mesure de le modifier.",
        "confirm" : "Envoyer",
        "cancel" : "Annuler"
      },
      "steps": {
        "saveAsDraft" : "Enregistrer brouillon",
        "previous": "Etape précédente",
        "next": "Etape suivante",
        "validate": "Valider",
        "close": "Fermer",
        "yes": "Oui",
        "no": "Non",
        "step": "Etape",
        "done": "Terminé !",
        "required": "* Champ requis."
      },
      "result": {
        "success": "Bravo !",
        "successMessage": "Votre présentation a été envoyée. Vous recevrez bientôt un email de confirmation. Nous vous recontacterons dès que nous aurons fait notre choix.",
        "goToHome": "Retour à la page principale"
      },
      "admin": {
        "logout": "Se déconnecter",
        "session": "Talk",
        "sessions": "Talks",
        "administration": "Administration",
        "toggle": "Ouvrir le volet",
        "clearSorting": "Annuler le tri",
        "clearFilter": "Annuler les filtres",
        "speaker": "Conférencier",
        "title": "Titre",
        "difficulty": "Difficulté",
        "track": "Track",
        "description": "Description",
        "mean": "Moyenne",
        "date": "Date",
        "deliberation": "Délibération",
        "commentaries": "Commentaires",
        "message": "Message",
        "votes": "Votes",
        "you": "Vous",
        "financialHelp": "Aide financière"
      },
      "config": {
        "logout": "Se déconnecter",
        "login": "Se connecter",
        "config": "Configuration",
        "linkMyAccount" : "Lier mon compte Google Drive avec Call For Paper",
        "configurationNeeded" : "L'administrateur doit configurer l'application avant que vous puissiez l'utiliser",
        "success" : "Votre compte est maintenant lié",
        "error" : "Error lors de la configuration",
        "configureLink" : "Configurer l'application"
      },
      "error": {
        "backendcommunication": "Désolé, il y a eu un problème avec le serveur distant",
        "noInternet": "Désolé, il y a eu une problème de connexion, êtes vous connecté à internet ?"
      },
      "just_now": "à l'instant",
      "seconds_ago": "il y a {{time}} secondes",
      "a_minute_ago": "il y a une minute",
      "minutes_ago": "il y a {{time}} minutes",
      "an_hour_ago": "il y a une heure",
      "hours_ago": "il y a {{time}} heures",
      "a_day_ago": "hier",
      "days_ago": "il y a {{time}} jours",
      "a_week_ago": "il y a une semaine",
      "weeks_ago": "il y a {{time}} semaines",
      "a_month_ago": "il y a un mois",
      "months_ago": "il y a {{time}} mois",
      "a_year_ago": "il y a un an",
      "years_ago": "il y a {{time}} ans",
      "over_a_year_ago": "il y a plus d'un an",
      "seconds_from_now": "dans une seconde",
      "a_minute_from_now": "dans une minute",
      "minutes_from_now": "dans {{time}} minutes",
      "an_hour_from_now": "dans une heure",
      "hours_from_now": "dans {{time}} heures",
      "a_day_from_now": "demain",
      "days_from_now": "dans {{time}} jours",
      "a_week_from_now": "dans une semaine",
      "weeks_from_now": "dans {{time}} semaine",
      "a_month_from_now": "dansun mois",
      "months_from_now": "dans {{time}} mois",
      "a_year_from_now": "dans un an",
      "years_from_now": "dans {{time}} ans",
      "over_a_year_from_now": "dans plus d'un an"
    });
    $translateProvider.translations('en-US', {
      "lang": {
        "fr-FR": "French",
        "en-US": "English"
      },
      "header": {
        "login": "Login",
        "logout": "Logout",
        "dashboard": "Dashboard",
        "profile": "Profile"
      },
      "dashboard" :
      {
        "createNewSession" : "Create a new talk",
        "editionTalks" : "Talks available for editing",
        "noEditionTalks" : "You don't have any draft",
        "sendedTalks" : "Submitted talks",
        "draftModifiedAt" : "Draft modified the ",
        "postedAt" : "submited the ",
        "noSendedTalks" : "You don't have submitted any talk yet",
        "verification" : "Verification",
        "verificationNeeded" : "An email has been sent to your address, click the link in it to confirm your account.",
      },
      "profile" :
      {
        "save" : "Save",
        "success" : "Profile saved",
        "error" : "Error saving the profile",
      },
      "login" : {
        "email" : "Email",
        "password" : "Password",
        "login" : "Login",
        "noAccount" : "Don't have an account yet?",
        "signup" : "Signup",
        "or" : "Or",
        "signGoogle" : "Sign in with Google",
        "signGithub" : "Sign in with Github",
        "wait": "Please wait",
        "badCredentials": "Incorrect login or password",
        "alreadyLinked" : "There is already a user associated with this  provider"
      },
      "signup" : {
        "signup" : "Sign up",
        "email" : "Email",
        "password" : "Password",
        "confirmPassword" : "Confirm Password",
        "yesAccount" : "Already have an account?",
        "loginNow" : "Log in now",
        "emailRequired" : "Your email address is required.",
        "emailPattern" : "Your email address is invalid.",
        "passwordMinLength" : "Password must be at least 6 characters long",
        "passwordRequired" : "Password is required.",
        "passwordMatch" : "Password must match.",
        "alreadyExists" : "There is already a user associated with this email"
      },
      "verify" : {
        "title" : "Email ",
        "verified" : "Your email has been verified",
        "alreadyVerified" : "This account is already verified",
        "notVerified" : "Errror during verification"
      },
      "step1": {
        "email": "Email *:",
        "name": "Name *:",
        "firstname": "Firstname *:",
        "phone": "Phone:",
        "errorPhone": "Please enter a correct phone number",
        "company": "Company:",
        "bio": "Bio *:",
        "hintBio": "Describe yourself with a few words. This description will be use to fill the website.",
        "social": "URL (Google, Github, etc):",
        "hintSocial": "Give us your socials networks data (for the website) : Twitter / G+ / Github / Blog."
      },
      "step2": {
        "name": "Session name *",
        "description": "Description *:",
        "hintDescription": "Give a description of your talk. This description will be used to fill the website.",
        "references": "References or complement informations :",
        "hintReferences": "Is there any conferences where you have already spoken ? If you could give a link to the presentation, it's better.",
        "complement": "Recommendation and additional information:",
        "difficulty": "Difficulty * (Beginner, Confirmed, Expert):",
        "track": "Track *:",
        "cospeaker": "Co-speaker:",
        "hintCospeaker": "If you are not alone on stage, give the co-speaker name / email / bio / social networks.",
        "beginner": "Beginner",
        "confirmed": "Confirmed",
        "expert": "Expert",
        "tracks": {
          "cloud": "Cloud",
          "mobile": "Internet of things",
          "web": "Web",
          "discovery": "Discovery"
        },
        "type": "Type *:",
        "hintType": "A conference must last about 45 minutes and a codelab about 2 hours.",
        "types": {
          "conference": "Conference",
          "codelab": "Codelab"
        },
        "hintTrack": "Choose the track where you think your talk will be place."
      },
      "step3": {
        "header1": "Here is all the informations relatives to your venue. The following informations will be carefuly study for our decision. So please select Yes, only if needed.",
        "header2": "Breakfast and lunch is offered the days of the event.",
        "financial": "Do you need financial help ? *:",
        "labelTravel": "Travel",
        "travel": "I need financial help for the trip.",
        "date": "For whitch date(s) ?:",
        "labelHotel": "Housing",
        "hotel": "I need financial help for the hotel.",
        "place": "Where are you coming from ?:",
        "sendError": "An error occurred during the submission, please retry."
      },
      "confirmModal": {
        "title" : "Confirmation",
        "text" : "Do you really want to send this talk ? Once sent you will not be able to modify it.",
        "confirm" : "OK",
        "cancel" : "Cancel"
      },
      "steps": {
        "saveAsDraft" : "Save as draft",
        "previous": "Previous step",
        "next": "Next step",
        "validate": "Submit",
        "close": "Close",
        "yes": "Yes",
        "no": "No",
        "step": "Step",
        "done": "Done !",
        "required": "* Required field."
      },
      "result": {
        "success": "Well done !",
        "successMessage": "Your talk has been send. You will soon receive a confirmation email. We will contact you as soon as we will make our decision.",
        "goToHome": "Back to main page"
      },
      "admin": {
        "logout": "Logout",
        "session": "Talk",
        "sessions": "Talks",
        "administration": "Administration",
        "toggle": "Toggle navigation",
        "clearSorting": "Clear sorting",
        "clearFilter": "Clear filters",
        "speaker": "Speaker",
        "title": "Title",
        "difficulty": "Difficulty",
        "track": "Track",
        "description": "Description",
        "mean": "Mean",
        "date": "Date",
        "deliberation": "Deliberation",
        "commentaries": "Commentaries",
        "message": "Message",
        "votes": "Votes",
        "you": "You",
        "financialHelp": "Financial Help"
      },
      "config": {
        "login": "Login",
        "logout": "Logout",
        "config": "Configuration",
        "linkMyAccount" : "Link my Google Drive account with Call For Paper",
        "configurationNeeded" : "Administrator must configure the application before you use it",
        "success" : "Your account has been linked",
        "error" : "Error linking your account with Google Drive",
        "configureLink" : "Configure the application"
      },
      "error": {
        "backendcommunication": "Sorry, a problem occure with the server",
        "noInternet": "Sorry, it seems that your are not connected to internet"
      },
      "just_now": "just now",
      "seconds_ago": "{{time}} seconds ago",
      "a_minute_ago": "a minute ago",
      "minutes_ago": "{{time}} minutes ago",
      "an_hour_ago": "an hour ago",
      "hours_ago": "{{time}} hours ago",
      "a_day_ago": "yesterday",
      "days_ago": "{{time}} days ago",
      "a_week_ago": "a week ago",
      "weeks_ago": "{{time}} weeks ago",
      "a_month_ago": "a month ago",
      "months_ago": "{{time}} months ago",
      "a_year_ago": "a year ago",
      "years_ago": "{{time}} years ago",
      "over_a_year_ago": "over a year ago",
      "seconds_from_now": "{{time}} seconds from now",
      "a_minute_from_now": "a minute from now",
      "minutes_from_now": "{{time}} minutes from now",
      "an_hour_from_now": "an hour from now",
      "hours_from_now": "{{time}} hours from now",
      "a_day_from_now": "tomorrow",
      "days_from_now": "{{time}} days from now",
      "a_week_from_now": "a week from now",
      "weeks_from_now": "{{time}} weeks from now",
      "a_month_from_now": "a month from now",
      "months_from_now": "{{time}} months from now",
      "a_year_from_now": "a year from now",
      "years_from_now": "{{time}} years from now",
      "over_a_year_from_now": "over a year from now"
    });
    $translateProvider.preferredLanguage('fr-FR');
  }]);