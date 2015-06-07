package cfp

import(
    "net/http"
    "encoding/json"
    "io/ioutil"
    "io"
	"appengine"
	"appengine/user"
	"appengine/datastore"
	"github.com/gorilla/context"

	"src/gorequest"
)

/**
 * set admin user context
 * @param {[type]} w    http.ResponseWriter [description]
 * @param {[type]} r    *http.Request       [description]
 * @param {[type]} next http.HandlerFunc    [description]
 */
func HandleAdmin(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	c := appengine.NewContext(r)

	u := user.Current(c)
	connected := (u != nil)
	if !connected {
		return
	} else {
		adminUser := AdminUser{
			Email : u.Email,
			Name : u.Email,
		}
		currentUserKey := datastore.NewKey(c, "AdminUser", u.ID, 0, nil)
		context.Set(r, "adminUser", adminUser)
		key, _ := datastore.Put(c, currentUserKey, &adminUser)
		context.Set(r, "adminUserKey", key)
	 	next(w, r)
	}
}

/**
 * get logout link with redirection url
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request       [description]
 * @return {[type]}   [description]
 */
func getAdminUserLogout(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }

	redirect := Redirect{}
	if err := json.Unmarshal(body, &redirect); err != nil {
        SetError(w, r, 400, "")
		return
    }

	u := user.Current(c)
	connected := (u != nil)
	if !connected {
	 	uriResponse := &Uri{
			Uri : "",
	    }
	    json.NewEncoder(w).Encode(uriResponse)
	} else {
		uri, _ := user.LoginURL(c, redirect.Redirect)
	 	uriResponse := &Uri{
			Uri : uri,
	    }
	    json.NewEncoder(w).Encode(uriResponse)
	}
	return
}

/**
 * get admin user info and connect/disconnect link
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request       [description]
 * @return {[type]}   [description]
 */
func getCurrentUser(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	u := user.Current(c)
	connected := (u != nil)
	if !connected {
	 	uri, _ := user.LoginURL(c, "/")
	 	currentUser := &CurrentUser{
			Connected : connected,
	 		Admin : false,
	 		Config : true,
	 		Uri : uri,
	    }
	    json.NewEncoder(w).Encode(currentUser)
	} else {
	 	uri, _ := user.LogoutURL(c, "/")
	 	currentUser := &CurrentUser{
			Connected : connected,
	 		Admin : u.Admin,
	 		Config : true,
	 		Uri : uri,
		}
	 	json.NewEncoder(w).Encode(currentUser)
	}
	return
}

/**
 * get login link with redirection link
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request       [description]
 * @return {[type]}   [description]
 */
func getAdminUserLogin(w http.ResponseWriter, r *http.Request) {
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

	redirect := Redirect{}
	if err := json.Unmarshal(body, &redirect); err != nil {
        SetError(w, r, 400, "")
		return
    }

	u := user.Current(c)
	connected := (u != nil)
	if !connected {
	 	uri, _ := user.LoginURL(c, redirect.Redirect)
	 	uriResponse := &Uri{
			Uri : uri,
	    }
	    json.NewEncoder(w).Encode(uriResponse)
	} else {
	 	uriResponse := &Uri{
			Uri : "",
	    }
	    json.NewEncoder(w).Encode(uriResponse)
	}
	return
}


/**
 * save notififcatin token for current logged in user
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request       [description]
 * @return {[type]}   [description]
 */
func postNotifToken(w http.ResponseWriter, r *http.Request) {
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

    token := Token{}
	if err := json.Unmarshal(body, &token); err != nil {
        SetError(w, r, 400, "")
		return
    }

	adminUser := AdminUser{}
	key := context.Get(r, "adminUserKey").(*datastore.Key)
	err = datastore.Get(c, key, &adminUser)
	if err != nil {
		// User doesn't exists
		// 404
        SetError(w, r, 404, "user not found")
		return
	}
	adminUser.Notif = token.Token
	datastore.Put(c, key, &adminUser)
	json.NewEncoder(w).Encode(true)
	return
}

func sendNotif(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	notif := Notif{}

	query := datastore.NewQuery("AdminUser")
	// Get all adminUsers
	adminUsers := []AdminUser{}
	keys, err := query.GetAll(c, &adminUsers)
	if err != nil {
		return
	}
	ids := []string{}
	// Add notif ids to array
	for i, _ := range keys {
		if adminUsers[i].Notif != "" {
			ids = append(ids, adminUsers[i].Notif)
		}
	}
    notif.Ids = ids;
    data, err := json.Marshal(notif)
    if err != nil {
        return
    }

	gorequest.New(r).Post("https://android.googleapis.com/gcm/send").
		Set("Authorization", "key=" + config.Get.NOTIF_SERVER_KEY).
		Send(string(data)).
		End()
}

func insert(original []string, position int, value string) []string {
  l := len(original)
  target := original
  if cap(original) == l {
    target = make([]string, l+1, l+10)
    copy(target, original[:position])
  }
  copy(target[position+1:], original[position:])
  target[position] = value
  return target
}