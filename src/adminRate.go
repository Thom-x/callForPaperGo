package cfp

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/gorilla/context"
    "encoding/json"
    "gopkg.in/validator.v2"
    "io"
    "io/ioutil"
	"appengine"
	"appengine/datastore"
	"strconv"
)


/**
 * post admin rate
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request       [description]
 * @return {[type]}   [description]
 */
func postRate(w http.ResponseWriter, r *http.Request) {
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

	postedAdminRate := AdminRate{}
	if err := json.Unmarshal(body, &postedAdminRate); err != nil {
        SetError(w, r, 400, "")
		return
    }

    if err := validator.Validate(postedAdminRate); err != nil {
        w.WriteHeader(400) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
        return
	}

	rowId, err := strconv.ParseInt(postedAdminRate.RowId,10,64)
	if err != nil {
        SetError(w, r, 400, "")
		return
	}
	key := datastore.NewKey(c, "AdminRate", "", rowId,context.Get(r,"adminUserKey").(*datastore.Key))
	key, err = datastore.Put(c, key, &postedAdminRate)
	if err != nil {
        SetError(w, r, 400, "")
		return
	}
	rateReponse := AdminRateReponse{
		ID : key.IntID(),
		Rate : postedAdminRate.Rate,
		User : context.Get(r, "adminUser").(AdminUser),
	}

	json.NewEncoder(w).Encode(rateReponse)
	return
}


/**
 * get connected admin rate giving the session 
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request       [description]
 * @return {[type]}   [description]
 */
func getAdminRateByRowIdAndUserMe(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	id := mux.Vars(r)["id"]

	query := datastore.NewQuery("AdminRate").Ancestor(context.Get(r,"adminUserKey").(*datastore.Key)).Filter("RowId =", id)
	// If the user is logged in fetch only their lists.
	adminRates := []AdminRate{}
	keys, err := query.GetAll(c, &adminRates)
	if err != nil {
        SetError(w, r, 400, "")
		return
	}

	if len(adminRates) == 0 {
		json.NewEncoder(w).Encode(AdminRateReponse{})
		return 
	}

	adminRatesResponse := AdminRateReponse{}

	// Update the encoded keys and encode the rates.
	for i, k := range keys {
		adminRates[i].ID = k.IntID()
		var adminUser AdminUser
		datastore.Get(c, k.Parent(), &adminUser)
		adminRatesResponse.ID = adminRates[i].ID
		adminRatesResponse.Rate = adminRates[i].Rate
		adminRatesResponse.RowId = id
		adminRatesResponse.User = adminUser
	}
	json.NewEncoder(w).Encode(adminRatesResponse)
	return
}


/**
 * get rates by session
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request       [description]
 * @return {[type]}   [description]
 */
func getAdminRateByRow(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	id := mux.Vars(r)["id"]

	query := datastore.NewQuery("AdminRate").Filter("RowId =", id)
	// If the user is logged in fetch only their lists.
	adminRates := []AdminRate{}
	keys, err := query.GetAll(c, &adminRates)
	if err != nil {
        SetError(w, r, 400, "")
		return
	}

	adminRatesResponses := make([]AdminRateReponse,len(keys))

	// Update the encoded keys and encode the rates.
	for i, k := range keys {
		adminRates[i].ID = k.IntID()
		var adminUser AdminUser
		datastore.Get(c, k.Parent(), &adminUser)
		adminRatesResponses[i].ID = adminRates[i].ID
		adminRatesResponses[i].Rate = adminRates[i].Rate
		adminRatesResponses[i].RowId = id
		adminRatesResponses[i].User = adminUser
	}
	json.NewEncoder(w).Encode(adminRatesResponses)
	return
}
