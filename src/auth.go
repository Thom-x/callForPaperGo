package cfp

import (
	"log"
    "golang.org/x/crypto/bcrypt"
    "net/http"
    "encoding/json"
    "strconv"
    "gopkg.in/validator.v2"
    "math/rand"
    "io/ioutil"
    "io"
	"appengine"
	"appengine/datastore"
)

/**
 * verify user account
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request       [description]
 * @return {[type]}   [description]
 */
func getVerify(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	params := r.URL.Query()
  	id := params.Get("id")
  	token := params.Get("token")

	query := datastore.NewQuery("User")
	// If the user is logged in fetch only their lists.
	users := []User{}
	keys, err := query.GetAll(c, &users)
	if err != nil {
		SetError(w, r, 400, "error verifying your account")
		return
	}

	// Update the encoded keys and encode the lists.
	user := User{}
	found := false
	var key *datastore.Key
	for i, k := range keys {
		if strconv.FormatInt(k.IntID(),10) == id {
			found = true
			user = users[i];
			user.ID = k.IntID()
			key = k
		}
	}

	if !found {
		SetError(w, r, 404, "user not found")
		return
	}

	if user.Verified == true {
		SetError(w, r, 409, "already verified")
		return
	}
	if user.Verified == false && user.TokenVerified == token {
		user.Verified = true;
		user.TokenVerified = "";
		key, _ = datastore.Put(c, key, &user)

		SetToken(w, r, &user)
	}
	return
}

/**
 * signup new user
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request       [description]
 * @return {[type]}   [description]
 */
func postSignup(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        SetError(w, r, 400, "")
		return
    }
    if err := r.Body.Close(); err != nil {
        SetError(w, r, 400, "")
		return
    }

	signUpUser := SignUpUser{}
	if err := json.Unmarshal(body, &signUpUser); err != nil {
        SetError(w, r, 400, "")
		return
    }

    if err := validator.Validate(signUpUser); err != nil {
        w.WriteHeader(400) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
        return
	}

	query := datastore.NewQuery("User").Filter("Email =", signUpUser.Email)
	// If the user is logged in fetch only their lists.
	users := []User{}
	_, err = query.GetAll(c, &users)
	if err != nil {
        SetError(w, r, 400, "")
		return
	}
	if len(users) > 0 {
		return
	}

    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(signUpUser.Password), bcrypt.DefaultCost )
	hashedPasswordString := string(hashedPassword[:])

	newUser := User{}
	newUser.Password = hashedPasswordString
	newUser.Email = signUpUser.Email
	newUser.Profile = "{}"
	newUser.Verified = false

	newUser.TokenVerified = randSeq(100)

	key := datastore.NewIncompleteKey(c, "User", userKey(c))
	key, err = datastore.Put(c, key, &newUser)
	if err != nil {
        SetError(w, r, 400, "")
		return
	}
	newUser.ID = key.IntID()

	link := "/#/verify?id=" + strconv.FormatInt(key.IntID(),10) + "&token=" + newUser.TokenVerified
	log.Println(link)
	SendEmail(r, "verify.html", map[string]string{"hostname": config.Get.HOSTNAME, "link" : link,}, "Verifier votre adresse e-mail", newUser.Email, config.Get.EMAIL_SENDER)
	SetToken(w, r, &newUser)
	return
}

/**
 * generate random token
 */
func randSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

/**
 * login user
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request       [description]
 * @return {[type]}   [description]
 */
func postLogin(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }

	loginUser := LoginUser{}
	if err := json.Unmarshal(body, &loginUser); err != nil {
        SetError(w, r, 400, "")
		return
    }

    if err := validator.Validate(loginUser); err != nil {
        SetError(w, r, 400, "")
		return
	}

	query := datastore.NewQuery("User").Filter("Email =", loginUser.Email)
	// If the user is logged in fetch only their lists.
	users := []User{}
	_, err = query.GetAll(c, &users)
	if err != nil {
        SetError(w, r, 400, "")
		return
	}
	if len(users) == 0 {
        SetError(w, r, 401, "bad credentials")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte(loginUser.Password))
    if err != nil{
    	SetError(w, r, 401, "bad credentials")
		return
    }

    SetToken(w, r, &users[0])
	return 
}
