package cfp

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/gorilla/context"
    "encoding/json"
    "gopkg.in/validator.v2"
    "io"
    "io/ioutil"
	"time"
	"appengine"
	"appengine/datastore"
)

/**
 * get admin comments by session
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request       [description]
 * @return {[type]}   [description]
 */
func getAdminCommentByRow(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	id := mux.Vars(r)["id"]

	query := datastore.NewQuery("AdminComment").Filter("RowId =", id)
	// If the user is logged in fetch only their lists.
	adminComments := []AdminComment{}
	keys, err := query.GetAll(c, &adminComments)
	if err != nil {
        SetError(w, r, 400, "")
		return
	}

	adminCommentResponses := make([]AdminCommentResponse,len(keys))

	// Update the encoded keys and encode the rates.
	for i, k := range keys {
		adminComments[i].ID = k.IntID()
		var adminUser AdminUser
		datastore.Get(c, k.Parent(), &adminUser)
		adminCommentResponses[i].ID = adminComments[i].ID
		adminCommentResponses[i].Comment = adminComments[i].Comment
		adminCommentResponses[i].RowId = id
		adminCommentResponses[i].Added = adminComments[i].Added
		adminCommentResponses[i].User = adminUser
	}
	json.NewEncoder(w).Encode(adminCommentResponses)
	return
}

/**
 * post comment
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request       [description]
 * @return {[type]}   [description]
 */
func postComment(w http.ResponseWriter, r *http.Request) {
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

	postedComment := AdminComment{}
	if err := json.Unmarshal(body, &postedComment); err != nil {
        SetError(w, r, 400, "")
		return
    }

    if err := validator.Validate(postedComment); err != nil {
        w.WriteHeader(400) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
        return
	}

    postedComment.Added = time.Now().UnixNano() / int64(time.Millisecond)

	key := datastore.NewIncompleteKey(c, "AdminComment", context.Get(r, "adminUserKey").(*datastore.Key))
	key, err = datastore.Put(c, key, &postedComment)
	if err != nil {
        SetError(w, r, 400, "")
		return
	}
	adminCommentResponse := AdminCommentResponse{
		ID : key.IntID(),
		Comment : postedComment.Comment,
		Added : postedComment.Added,
		RowId : postedComment.RowId,
		User : context.Get(r, "adminUser").(AdminUser),
	}

	json.NewEncoder(w).Encode(adminCommentResponse)
	return
}