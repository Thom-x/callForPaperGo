package cfp

import(
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
    "strconv"
    "gopkg.in/validator.v2"
    "io/ioutil"
    "io"
	"time"
	"appengine"
	"appengine/datastore"
)

/**
 * get current user drafts
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request       [description]
 * @return {[type]}   [description]
 */
func getDrafts(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	currentUserKey, _, _ := getCurrentUserKey(c,r)

	query := datastore.NewQuery("Session").Ancestor(currentUserKey).Filter("Draft =", true)
	// If the user is logged in fetch only their lists.
	sessions := []Session{}
	keys, err := query.GetAll(c, &sessions)
	if err != nil {
        SetError(w, r, 400, "")
		return
	}

	// Update the encoded keys and encode the lists.
	for i, k := range keys {
		sessions[i].ID = k.IntID()
	}
	json.NewEncoder(w).Encode(sessions)
	return 
}

/**
 * get draft
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request       [description]
 * @return {[type]}   [description]
 */
func getDraft(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	//Get the sessions id from the URL.
	id := mux.Vars(r)["id"]

	currentUserKey, _,_  := getCurrentUserKey(c,r)

	query := datastore.NewQuery("Session").Ancestor(currentUserKey).Filter("Draft =", true)
	// If the user is logged in fetch only their lists.
	sessions := []Session{}
	keys, err := query.GetAll(c, &sessions)
	if err != nil {
        SetError(w, r, 400, "")
		return
	}

	// Update the encoded keys and encode the lists.
	session := Session{}
	found := false
	for i, k := range keys {
		if strconv.FormatInt(k.IntID(),10) == id {
			found = true
			session = sessions[i];
			session.ID = k.IntID()
		}
	}
	if !found {
        SetError(w, r, 404, "draft not found")
		return
	}
	json.NewEncoder(w).Encode(session)
	return
}

/**
 * delete draft
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request       [description]
 * @return {[type]}   [description]
 */
func deleteDraft(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	//Get the sessions id from the URL.
	id := mux.Vars(r)["id"]

	currentUserKey, _, _ := getCurrentUserKey(c,r)

	query := datastore.NewQuery("Session").Ancestor(currentUserKey).Filter("Draft =", true)
	// If the user is logged in fetch only their lists.
	sessions := []Session{}
	keys, err := query.GetAll(c, &sessions)
	if err != nil {
        SetError(w, r, 400, "")
		return
	}

	// Update the encoded keys and encode the lists.
	session := Session{}
	found := false
	for _, k := range keys {
		if strconv.FormatInt(k.IntID(),10) == id {
			datastore.Delete(c, k)
			found = true
		}
	}
	if !found {
        SetError(w, r, 404, "draft not found")
		return
	}
	json.NewEncoder(w).Encode(session)
	return
}

/**
 * post draft
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request       [description]
 * @return {[type]}   [description]
 */
func postDraft(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }

	session := Draft{}
	if err := json.Unmarshal(body, &session); err != nil {
        SetError(w, r, 400, "")
		return
    }
    
    if err := validator.Validate(session); err != nil {
        w.WriteHeader(400) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
        return
	}
    session.Added = time.Now().UnixNano() / int64(time.Millisecond)
    session.Draft = true

    currentUserKey, _, _ := getCurrentUserKey(c,r)
	key := datastore.NewIncompleteKey(c, "Session", currentUserKey)
	key, err = datastore.Put(c, key, &session)
	if err != nil {
        SetError(w, r, 400, "")
		return
	}
	session.ID = key.IntID()

	json.NewEncoder(w).Encode(session)
	return
}

/**
 * put draft
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request       [description]
 * @return {[type]}   [description]
 */
func putDraft(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	//Get the sessions id from the URL.
	id := mux.Vars(r)["id"]

	currentUserKey, _, _ := getCurrentUserKey(c,r)

	query := datastore.NewQuery("Session").Ancestor(currentUserKey).Filter("Draft =", true)
	// If the user is logged in fetch only their lists.
	sessions := []Session{}
	keys, err := query.GetAll(c, &sessions)
	if err != nil {
        SetError(w, r, 400, "")
		return
	}

	// Update the encoded keys and encode the lists.
	session := Session{}
	found := false
	var key *datastore.Key
	for i, k := range keys {
		if strconv.FormatInt(k.IntID(),10) == id {
			found = true
			session = sessions[i];
			session.ID = k.IntID()
			key = k
		}
	}
	if !found {
        SetError(w, r, 404, "draft not found")
		return
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        SetError(w, r, 400, "")
		return
    }
    if err := r.Body.Close(); err != nil {
        SetError(w, r, 400, "")
		return
    }

	postedSession := Draft{}
	if err := json.Unmarshal(body, &postedSession); err != nil {
        SetError(w, r, 400, "")
		return
    }

    if err := validator.Validate(postedSession); err != nil {
        w.WriteHeader(400) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
        return
	}

    postedSession.Added = time.Now().UnixNano() / int64(time.Millisecond)
    postedSession.Draft = true
	key, err = datastore.Put(c, key, &postedSession)
	if err != nil {
        SetError(w, r, 400, "")
		return
	}
	session.ID = key.IntID()

	json.NewEncoder(w).Encode(postedSession)
	return
}

/**
 * get user sessions
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request       [description]
 * @return {[type]}   [description]
 */
func getSessions(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	currentUserKey, _, _ := getCurrentUserKey(c,r)

	query := datastore.NewQuery("Session").Ancestor(currentUserKey).Filter("Draft =", false)
	// If the user is logged in fetch only their lists.
	sessions := []Session{}
	keys, err := query.GetAll(c, &sessions)
	if err != nil {
        SetError(w, r, 400, "")
		return
	}

	// Update the encoded keys and encode the lists.
	for i, k := range keys {
		sessions[i].ID = k.IntID()
	}
	json.NewEncoder(w).Encode(sessions)
	return 
}

/**
 * put session
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request       [description]
 * @return {[type]}   [description]
 */
func putSession(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	//Get the sessions id from the URL.
	id := mux.Vars(r)["id"]

	currentUserKey, _, _ := getCurrentUserKey(c,r)

	query := datastore.NewQuery("Session").Ancestor(currentUserKey).Filter("Draft =", true)
	// If the user is logged in fetch only their lists.
	sessions := []Session{}
	keys, err := query.GetAll(c, &sessions)
	if err != nil {
        SetError(w, r, 400, "")
		return
	}

	// Update the encoded keys and encode the lists.
	session := Session{}
	found := false
	var key *datastore.Key
	for i, k := range keys {
		if strconv.FormatInt(k.IntID(),10) == id {
			found = true
			session = sessions[i];
			session.ID = k.IntID()
			key = k
		}
	}
	if !found {
        SetError(w, r, 404, "session not found")
		return
	}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        SetError(w, r, 400, "")
		return
    }
    if err := r.Body.Close(); err != nil {
        SetError(w, r, 400, "")
		return
    }

	postedSession := Session{}
	if err := json.Unmarshal(body, &postedSession); err != nil {
        SetError(w, r, 400, "")
		return
    }

    if err := validator.Validate(postedSession); err != nil {
        w.WriteHeader(400) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
        return
	}

    postedSession.Added = time.Now().UnixNano() / int64(time.Millisecond)
    postedSession.Draft = false
	key, err = datastore.Put(c, key, &postedSession)
	if err != nil {
        SetError(w, r, 400, "")
		return
	}
	session.ID = key.IntID()

	SendEmail(r, "confirmed.html", map[string]string{"hostname": config.Get.HOSTNAME, "name" : session.Firstname, "talk" : session.SessionName}, "Votre session a bien été enregistré", session.Email, config.Get.EMAIL_SENDER)
	sendNotif(w, r)
	json.NewEncoder(w).Encode(postedSession)
	return
}

/**
 * get session
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request       [description]
 * @return {[type]}   [description]
 */
func getSession(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	//Get the sessions id from the URL.
	id := mux.Vars(r)["id"]

	currentUserKey, _, _ := getCurrentUserKey(c,r)

	query := datastore.NewQuery("Session").Ancestor(currentUserKey).Filter("Draft =", false)
	// If the user is logged in fetch only their lists.
	sessions := []Session{}
	keys, err := query.GetAll(c, &sessions)
	if err != nil {
        SetError(w, r, 400, "")
		return
	}

	// Update the encoded keys and encode the lists.
	session := Session{}
	found := false
	for i, k := range keys {
		if strconv.FormatInt(k.IntID(),10) == id {
			found = true
			session = sessions[i];
			session.ID = k.IntID()
		}
	}
	if !found {
        SetError(w, r, 404, "session not found")
		return
	}
	json.NewEncoder(w).Encode(session)
	return
}

/**
 * post session
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request       [description]
 * @return {[type]}   [description]
 */
func postSession(w http.ResponseWriter, r *http.Request) {
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

	session := Session{}
	if err := json.Unmarshal(body, &session); err != nil {
        SetError(w, r, 400, "")
		return
    }

    if err := validator.Validate(session); err != nil {
        w.WriteHeader(400) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
        return
	}

    session.Added = time.Now().UnixNano() / int64(time.Millisecond)
    session.Draft = false
    currentUserKey, _, _ := getCurrentUserKey(c,r)
	key := datastore.NewIncompleteKey(c, "Session", currentUserKey)
	key, err = datastore.Put(c, key, &session)
	if err != nil {
        SetError(w, r, 400, "")
		return
	}
	session.ID = key.IntID()

	SendEmail(r, "confirmed.html", map[string]string{"hostname": config.Get.HOSTNAME, "name" : session.Firstname, "talk" : session.SessionName}, "Votre session a bien été enregistré", session.Email, config.Get.EMAIL_SENDER)
	sendNotif(w, r)
	json.NewEncoder(w).Encode(session)
	return
}
