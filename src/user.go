package cfp

import(
    "net/http"
    "encoding/json"
	"appengine"
	"appengine/datastore"

	"gopkg.in/validator.v2"
    "io/ioutil"
    "io"
)

/**
 * return "default" user key
 * @param  {[type]} c appengine.Context [description]
 * @return {[type]}   [description]
 */
func userKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "User", "default", 0, nil)
}

/**
 * get user profile
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request       [description]
 * @return {[type]}   [description]
 */
func getUser(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	key, _, _ := getCurrentUserKey(c,r)
	var user User
	err := datastore.Get(c, key, &user)
	if err != nil {
		panic(err)
		// User doesn't exists
		// 404
        SetError(w, r, 404, "user not found")
		return
	}
	user.ID = key.IntID()

	profile := Profile{}
	if err := json.Unmarshal([]byte(user.Profile), &profile); err != nil {
        SetError(w, r, 400, "")
		return
    }
    profile.Email = user.Email;
	json.NewEncoder(w).Encode(profile)
	return
}

/**
 * save user profile
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request       [description]
 * @return {[type]}   [description]
 */
func putUser(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	key, _, _  := getCurrentUserKey(c,r)
	var user User
	err := datastore.Get(c, key, &user)
	if err != nil {
		// User doesn't exists
		// 404
        SetError(w, r, 404, "user not found")
		return
	}
	user.ID = key.IntID()

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        SetError(w, r, 400, "")
		return
    }
    if err := r.Body.Close(); err != nil {
        SetError(w, r, 400, "")
		return
    }

	profile := Profile{}
	if err := json.Unmarshal(body, &profile); err != nil {
        SetError(w, r, 400, "")
		return
    }

    if err := validator.Validate(profile); err != nil {
        w.WriteHeader(400) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
        return
	}

    profileJson, _ := json.Marshal(profile)
    user.Profile = string(profileJson[:])

	key, err = datastore.Put(c, key, &user)
	if err != nil {
        SetError(w, r, 400, "")
		return
	}
	return 
}
