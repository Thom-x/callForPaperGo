package cfp

import(
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
    "strconv"
	"appengine"
	"appengine/datastore"
)


/**
 * get session
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request       [description]
 * @return {[type]}   [description]
 */
func getAdminSession(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	//Get the sessions id from the URL.
	id := mux.Vars(r)["id"]

	query := datastore.NewQuery("Session").Filter("Draft =", false)
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
 * get all sessions
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request       [description]
 * @return {[type]}   [description]
 */
func getAdminSessions(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	query := datastore.NewQuery("Session").Filter("Draft =", false)
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
