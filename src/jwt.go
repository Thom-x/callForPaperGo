package cfp

import(
    "net/http"
    "github.com/gorilla/context"
    "github.com/dgrijalva/jwt-go"
    "strconv"
    "time"
    "encoding/json"
    "appengine"
    "appengine/datastore"
)

/**
 * get current user key using context
 * @param  {[type]} c appengine.Context [description]
 * @param  {[type]} r *http.Request)    (*datastore.Key, string, int64 [description]
 * @return {[type]}   [description]
 */
func getCurrentUserKey(c appengine.Context, r *http.Request) (*datastore.Key, string, int64) {
	user := context.Get(r, "user")
	token, _ := user.(*jwt.Token)
    sub, _ := token.Claims["sub"]
    currentUserIdStr := sub.(string)
	currentUserIdInt, _ := strconv.ParseInt(currentUserIdStr,10,64)
	key := datastore.NewKey(c, "User", "", currentUserIdInt, userKey(c))
	return key, currentUserIdStr, currentUserIdInt
}

/**
 * send jwt token to client
 * @param {[type]} w    http.ResponseWriter [description]
 * @param {[type]} r    *http.Request       [description]
 * @param {[type]} user *User               [description]
 */
func SetToken(w http.ResponseWriter, r *http.Request, user *User) {
	token := jwt.New(jwt.SigningMethodHS256)
    // Set some claims
    token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
    token.Claims["sub"] = strconv.FormatInt(user.ID,10)
    token.Claims["verified"] = user.Verified
    // Sign and get the complete encoded token as a string
    tokenString, _ := token.SignedString([]byte(config.Get.JWT_SECRET))
    tokenResponse := &Token{
    	Token: tokenString,
    }
    json.NewEncoder(w).Encode(tokenResponse)
	return
}
