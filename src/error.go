package cfp

import(
	"encoding/json"
	"net/http"
)

/**
 * throw error given http code and message
 * @param {[type]} w       http.ResponseWriter [description]
 * @param {[type]} r       *http.Request       [description]
 * @param {[type]} code    int                 [description]
 * @param {[type]} message string              [description]
 */
func SetError(w http.ResponseWriter, r *http.Request, code int, message string) {
    errorResponse := &Error{
    	Code : code,
    	Message : message,
    }
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(errorResponse)
	return
}

type Error struct {
    Message string `json:"message"`
    Code int `json:"code"`
}