
# Call For Paper Go

<img src="https://raw.githubusercontent.com/Thom-x/callForPaperGo/master/readme/screenshot.png" alt="alt text" width="100%">

## Features

 - Powered by App Engine
 - Golang Back-end
 - AngularJS Front-end
 - Datastore database
 - JWT Token
 - Localization (fr, en)
 - Material design
 - Google Cloud Messaging notification

### User Panel

 - Google, Github OAuth authentification
 - Account register (with email confirmation)
 - Talks submission (with email confirmation)
 - Save/delete draft

### Admin Panel

 - Google Cloud Messaging notifications
 - Sort talks (by rate, date, track...)
 - Filter talks (by track, talker name, description...)
 - Rate talk
 - Comment talk

## Setup

*Change **127.0.0.1:8080** to your application domain for production*

### Obtaining OAuth Keys

<img src="http://images.google.com/intl/en_ALL/images/srpr/logo6w.png" width="150">

- Visit [Google Cloud Console](https://cloud.google.com/console/project)
- Click **CREATE PROJECT** button
- Enter *Project Name*, then click **CREATE**
- Then select *APIs & auth* from the sidebar and click on *Credentials* tab
- Click **CREATE NEW CLIENT ID** button
 - **Application Type**: Web Application
 - **Authorized Javascript origins**: *http://127.0.0.1:8080*
 - **Authorized redirect URI**: *http://127.0.0.1:8080*

**Note:** Make sure you have turned on **Contacts API** and **Google+ API** in the *APIs* tab.

<hr>

<img src="https://www.cloudamqp.com/images/blog/github.png" height="70">

- Visit [Github developers settings](https://github.com/settings/developers)
- Click **REGISTER NEW APPLICATION** button
 - **Application name**: *Application name*
 - **Homepage URL**: *http://127.0.0.1:8080*
 - **Authorization callback URL**: *http://127.0.0.1:8080*

### Obtaining reCAPTCHA Keys

<img src="https://www.gstatic.com/recaptcha/admin/logo_recaptcha_color_24dp.png" height="70">

- Visit [reCaptcha panel](https://www.google.com/recaptcha/admin)
 - **Domaines**: *127.0.0.1:8080*
- Click **Save** button

### Obtaining GCM Keys

<img src="http://images.google.com/intl/en_ALL/images/srpr/logo6w.png" width="150">
- Select your project
- Then select *APIs & auth* from the sidebar and click on *Credentials* tab
- Click **CREATE NEW Key** button
 - **Server Key**
- Again click **CREATE NEW Key** button
 - **Browser Key**

**Note:** Make sure you have turned on **Contacts API** and **Google+ API** in the *APIs* tab.


### Edit CFP Settings


Edit `config.json` replace the informations to suit your need :

```json
{
    "GOOGLE_SECRET" : "yourGoogleClientSecret",
    "GITHUB_SECRET" : "yourGithubClientSecret",
    "EMAIL_SENDER" : "senderEmail",
    "JWT_SECRET" : "randomString",
    "NOTIF_SERVER_KEY" : "yourGCMApiServerKey",
    "EVENT_NAME" : "eventName",
    "COMMUNITY" : "community",
    "DATE" : "eventDate",
    "RELEASE_DATE" : "talksReleaseDate",
    "HOSTNAME" : "http://your-app-id.appspot.com"
}
```
Edit `app.yaml` replace the informations to suit your need :

```yaml
application: your-app-id
version: 1
runtime: go
api_version: go1

handlers:
- url: /
  static_files: static/dist/index.html
  upload: static/dist/index\.html

- url: /favicon\.ico
  static_files: static/dist/favicon.ico
  upload: static/dist/favicon\.ico

- url: /api/.*
  script: _go_app

- url: /auth/.*
  script: _go_app

- url: /*
  static_dir: static/dist

skip_files:
- ^(static/app/.*)
- ^(static/node_modules/.*)
- ^(node_modules/.*)
```
Edit `static/app/scripts/app.js` add your providers tokens :

```javascript
  .constant('Config', {
    'recaptcha': 'yourRecaptchaPublicToken',
    'googleClientId': 'yourGoogleClientId',
    'githubClientId': 'yourGithubClientId'
    'gcmApiKey': 'yourGCMApiBrowserKey'
  })
```
Edit `static/app/manifest.json` add your gcm_sender_id:

```javascript
  "gcm_sender_id": "projectid",
```

## Deployment :

### App Engine :

```shell
grunt build
goapp deploy .
```
 - Go to : http://YOUR_APP_ID.appspot.com


### Local :

```shell
grunt build
goapp serve .
```
 - Go to : http://127.0.0.1:8080 and http://127.0.0.1:8000 for App Engine panel

## Usage :

### Manage admin users :

- Visit [Google Cloud Console](https://cloud.google.com/console/project)
- Select your project
- Click **Permissions** button
- Click **Add member** button :
 - **Email**: New member e-mail
 - **Can view**
 
### App entry points :

 - http://127.0.0.1:8080/ : User login page (create new talks)

 - http://127.0.0.1:8080/#/admin : Admin panel (rating, comment...)