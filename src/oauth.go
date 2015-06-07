package cfp

import(
	"net/http"
    "github.com/gorilla/context"
	// "github.com/parnurzeal/gorequest"
	"src/gorequest"
    "encoding/json"
    "strconv"
	"appengine"
	"appengine/datastore"
	"github.com/google/go-querystring/query"
	"fmt"
	"net/url"
)

/**
 * OAuth 2
 */

func ServeJSON(w http.ResponseWriter, r *http.Request, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	js, err := json.Marshal(data)
  	if err != nil {
        SetError(w, r, 400, "")
		return
  	}
  	w.Write(js)
}


func newGoogleParams() *OAuth2Params {
	return &OAuth2Params{
		ClientSecret: config.Get.GOOGLE_SECRET,
		GrantType:    "authorization_code",
	}
}

func newGithubParams() *OAuth2Params {
	return &OAuth2Params{
		ClientSecret: config.Get.GITHUB_SECRET,
		GrantType:    "authorization_code",
	}
}

type Response map[string]interface{}

/**
 * Google Login
 */

/**
 * handler google callback login
 * @param {[type]} w http.ResponseWriter [description]
 * @param {[type]} r *http.Request       [description]
 */
func LoginWithGoogle(w http.ResponseWriter, r *http.Request) {

	accessTokenUrl := "https://accounts.google.com/o/oauth2/token"
	peopleApiUrl := "https://www.googleapis.com"
	peopleApiPath := "/plus/v1/people/me/openIdConnect"

	// Step 1. Exchange authorization code for access token.
	googleParams := newGoogleParams()
	googleParams.LoadFromHTTPRequest(r)

	v, _ := query.Values(googleParams)

	res, body, _ := gorequest.New(r).Post(accessTokenUrl).
		Send(v.Encode()).
		Type("form").
		End()

	if res.StatusCode != 200 {
		var errorData map[string]interface{}
		json.Unmarshal([]byte(body), &errorData)

		ServeJSON(w, r, &Response{
			"message": errorData["error"].(string),
		}, 500)
		return
	}

	// Step 2. Retrieve profile information about the current user.
	var atData accessTokenData
	json.Unmarshal([]byte(body), &atData)

	qs, _ := query.Values(atData)

	u, _ := url.ParseRequestURI(peopleApiUrl)
	u.Path = peopleApiPath
	u.RawQuery = qs.Encode()
	urlStr := fmt.Sprintf("%v", u)

	resProfile, body, _ := gorequest.New(r).Get(urlStr).End()


	var profileData map[string]interface{}
	json.Unmarshal([]byte(body), &profileData)

	if resProfile.StatusCode != 200 {
		ServeJSON(w, r, &Response{
			"message": profileData["error"].(map[string]interface{})["message"],
		}, 500)
		return
	}

	if IsTokenSet(w, r) {
		c := appengine.NewContext(r)

		_, currentUserId, _ := getCurrentUserKey(c,r)

		query := datastore.NewQuery("User")
		// If the user is logged in fetch only their lists.
		users := []User{}
		keys, err := query.GetAll(c, &users)
		if err != nil {
	        SetError(w, r, 400, "")
			return
		}
		user := User{}
		for i, k := range keys {
			if strconv.FormatInt(k.IntID(),10) == currentUserId && users[i].Google == profileData["sub"].(string){
				user = users[i];
				user.ID = k.IntID()
				// login
				SetToken(w, r, &user)
				return
			}
			if strconv.FormatInt(k.IntID(),10) != currentUserId && users[i].Google == profileData["sub"].(string){
				// Allready exists
		        SetError(w, r, 409, "account already linked")
				return
			}
		}

		// Link account
		currentUserIdInt, err := strconv.ParseInt(currentUserId,10,64)
		key := datastore.NewKey(c, "User", "", currentUserIdInt, userKey(c))
		user = User{}
		err = datastore.Get(c, key, &user)
		if err != nil {
			// User doesn't exists
			// 404
	        SetError(w, r, 404, "user not found")
			return
		}
		user.ID = key.IntID()
		user.Google = profileData["sub"].(string)
		key, err = datastore.Put(c, key, &user)
		if err != nil {
	        SetError(w, r, 400, "")
			return
		}
		SetToken(w, r, &user)

	} else {
		// Step 3b. Create a new user account or return an existing one.
		c := appengine.NewContext(r)

		query := datastore.NewQuery("User").Filter("Email =", profileData["email"].(string))
		// If the user is logged in fetch only their lists.
		users := []User{}
		keys, err := query.GetAll(c, &users)
		if err != nil {
	        SetError(w, r, 400, "")
			return
		}
		if len(users) > 0 {
			if users[0].Google != profileData["sub"].(string) {
		        SetError(w, r, 409, "account email already taken")
				return
			} else {
				// login
				users[0].ID = keys[0].IntID()
				SetToken(w, r, &users[0])
				return
			}
		}

		newUser := User{}
		newUser.Email = profileData["email"].(string)

		profile := Profile{}
		profile.Firstname = profileData["given_name"].(string)
		profile.Name = profileData["family_name"].(string)
		profileJson, _ := json.Marshal(profile)
    	newUser.Profile = string(profileJson[:])

		newUser.Verified = true
		newUser.Google = profileData["sub"].(string)

		key := datastore.NewIncompleteKey(c, "User", userKey(c))
		key, err = datastore.Put(c, key, &newUser)
		if err != nil {
	        SetError(w, r, 400, "")
			return
		}
		newUser.ID = key.IntID()

		SetToken(w, r, &newUser)
	}
}
/**
 * Github Login
 */

/**
 * handle github login callback
 * @param {[type]} w http.ResponseWriter [description]
 * @param {[type]} r *http.Request       [description]
 */
func LoginWithGithub(w http.ResponseWriter, r *http.Request) {
	accessTokenUrl := "https://github.com/login/oauth/access_token"
	peopleApiUrl := "https://api.github.com/user"
	peopleApiEmailsUrl := "https://api.github.com/user/emails"

	// Step 1. Exchange authorization code for access token.
	githubParams := newGithubParams()
	githubParams.LoadFromHTTPRequest(r)

	v, _ := query.Values(githubParams)

	res, body, _ := gorequest.New(r).Post(accessTokenUrl).
		Set("Accept", "application/json").
		Send(v.Encode()).
		Type("form").
		End()
		

	if res.StatusCode != 200 {
		var errorData map[string]interface{}
		json.Unmarshal([]byte(body), &errorData)

		ServeJSON(w, r, &Response{
			"message": errorData["error"].(string),
		}, 500)
		return
	}

	// Step 2. Retrieve profile information about the current user.
	var atData accessTokenData
	json.Unmarshal([]byte(body), &atData)

	qs, _ := query.Values(atData)

	u, _ := url.ParseRequestURI(peopleApiUrl)
	//u.Path = peopleApiPath
	u.RawQuery = qs.Encode()
	urlStr := fmt.Sprintf("%v", u)

	resProfile, body, _ := gorequest.New(r).Get(urlStr).End()

	var profileData map[string]interface{}
	json.Unmarshal([]byte(body), &profileData)

	if resProfile.StatusCode != 200 {
		ServeJSON(w, r, &Response{
			"message": profileData["error"].(map[string]interface{})["message"],
		}, 500)
		return
	}
	profileData["sub"] = strconv.FormatFloat(profileData["id"].(float64), 'f', 0, 64)

	u, _ = url.ParseRequestURI(peopleApiEmailsUrl)
	//u.Path = peopleApiPath
	u.RawQuery = qs.Encode()
	urlStr = fmt.Sprintf("%v", u)

	resProfile, body, _ = gorequest.New(r).Get(urlStr).End()

	var emailData []interface{}
	var emailErrorData map[string]interface{}
	json.Unmarshal([]byte(body), &emailData)

	if resProfile.StatusCode != 200 {
		ServeJSON(w, r, &Response{
			"message": emailErrorData["error"].(map[string]interface{})["message"],
		}, 500)
		return
	}

	for _,element := range emailData {
		field := element.(map[string]interface{})
		if field["primary"] == true {
			profileData["email"] = field["email"]
		}
	}

	if IsTokenSet(w, r) {
		c := appengine.NewContext(r)

		_, currentUserId, _ := getCurrentUserKey(c,r)

		query := datastore.NewQuery("User")
		// If the user is logged in fetch only their lists.
		users := []User{}
		keys, err := query.GetAll(c, &users)
		if err != nil {
	        SetError(w, r, 400, "")
			return
		}
		user := User{}
		for i, k := range keys {
			if strconv.FormatInt(k.IntID(),10) == currentUserId && users[i].Github == profileData["sub"].(string){
				user = users[i];
				user.ID = k.IntID()
				// login
				SetToken(w, r, &user)
				return
			}
			if strconv.FormatInt(k.IntID(),10) != currentUserId && users[i].Github == profileData["sub"].(string){
				// Allready exists
		        SetError(w, r, 400, "account already linked")
				return
			}
		}

		// Link account
		currentUserIdInt, err := strconv.ParseInt(currentUserId,10,64)
		key := datastore.NewKey(c, "User", "", currentUserIdInt, userKey(c))
		user = User{}
		err = datastore.Get(c, key, &user)
		if err != nil {
			// User doesn't exists
			// 404
	        SetError(w, r, 404, "user not found")
			return
		}
		user.ID = key.IntID()
		user.Github = profileData["sub"].(string)
		key, err = datastore.Put(c, key, &user)
		if err != nil {
	        SetError(w, r, 400, "")
			return
		}
		SetToken(w, r, &user)

	} else {
		// Step 3b. Create a new user account or return an existing one.
		c := appengine.NewContext(r)

		query := datastore.NewQuery("User").Filter("Email =", profileData["email"].(string))
		// If the user is logged in fetch only their lists.
		users := []User{}
		keys, err := query.GetAll(c, &users)
		if err != nil {
			SetError(w, r, 400, "")
			return
		}
		if len(users) > 0 {
			if users[0].Github != profileData["sub"].(string) {
				// Account with this email allready exists
				SetError(w, r, 409, "email account already taken")
				return
			} else {
				// login
				users[0].ID = keys[0].IntID()
				SetToken(w, r, &users[0])
				return
			}
		}

		newUser := User{}
		newUser.Email = profileData["email"].(string)

		profile := Profile{}
		profile.Company = profileData["company"].(string)
		bio := profileData["bio"]
		if bio != nil {
			profile.Bio = profileData["bio"].(string)
		}
		profile.Social = profileData["html_url"].(string)
		profileJson, _ := json.Marshal(profile)
    	newUser.Profile = string(profileJson[:])

		newUser.Verified = true
		newUser.Github = profileData["sub"].(string)

		key := datastore.NewIncompleteKey(c, "User", userKey(c))
		key, err = datastore.Put(c, key, &newUser)
		if err != nil {
	        SetError(w, r, 400, "")
			return
		}
		newUser.ID = key.IntID()

		SetToken(w, r, &newUser)
	}
}

/**
 * Utils
 */

func (f *OAuth2Params) LoadFromHTTPRequest(r *http.Request) {
	type requestData struct {
		Code        string `json:"code"`
		ClientId    string `json:"clientId"`
		RedirectUri string `json:"redirectUri"`
	}
	decoder := json.NewDecoder(r.Body)

	var data requestData
	err := decoder.Decode(&data)

	if err != nil {
		panic(err)
		return
	}
	f.Code = data.Code
	f.ClientId = data.ClientId
	f.RedirectUri = data.RedirectUri
}


func IsTokenSet(w http.ResponseWriter, r *http.Request) bool {
	_, ok := context.GetOk(r, "user")
	return ok
}

func NewClient() *http.Client {
	return &http.Client{}
}