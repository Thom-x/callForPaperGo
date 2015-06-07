package cfp

import (
    "net/http"
    "github.com/gorilla/mux"
    // "github.com/auth0/go-jwt-middleware"
    "src/go-jwt-middleware"
	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
    "encoding/json"
    "io/ioutil"
	"appengine"
	"appengine/mail"
	"text/template"
    "bytes"
)

var config Config

func init() {
	config.New()
	r := mux.NewRouter()

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Get.JWT_SECRET), nil
		},
	})

	jwtMiddlewareCommon := jwtmiddleware.New(jwtmiddleware.Options{
		ErrorHandler : nil,
		CredentialsOptional : true,
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Get.JWT_SECRET), nil
		},
	})

	auth := mux.NewRouter().PathPrefix("/auth").Subrouter()
	r.PathPrefix("/auth").Handler(negroni.New(
		negroni.HandlerFunc(jwtMiddlewareCommon.HandlerWithNextNoError),
		negroni.Wrap(auth),
	))
	auth.HandleFunc("/google", LoginWithGoogle).Methods("POST")
	auth.HandleFunc("/github", LoginWithGithub).Methods("POST")
	auth.HandleFunc("/signup", postSignup).Methods("POST")
	auth.HandleFunc("/login", postLogin).Methods("POST")
	auth.HandleFunc("/verify", getVerify).Methods("GET").Queries("id", "", "token", "")


	restricted := mux.NewRouter().PathPrefix("/api/restricted").Subrouter()
	r.PathPrefix("/api/restricted").Handler(negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(restricted),
	))
	restricted.HandleFunc("/draft", getDrafts).Methods("GET")
	restricted.HandleFunc("/draft", postDraft).Methods("POST")
	restricted.HandleFunc("/draft/{id}", getDraft).Methods("GET")
	restricted.HandleFunc("/draft/{id}", putDraft).Methods("PUT")
	restricted.HandleFunc("/draft/{id}", deleteDraft).Methods("DELETE")
	restricted.HandleFunc("/session", getSessions).Methods("GET")
	restricted.HandleFunc("/session", postSession).Methods("POST")
	restricted.HandleFunc("/session/{id}", getSession).Methods("GET")
	restricted.HandleFunc("/session/{id}", putSession).Methods("PUT")
	restricted.HandleFunc("/user", getUser).Methods("GET")
	restricted.HandleFunc("/user", putUser).Methods("PUT")

	admin := mux.NewRouter().PathPrefix("/api/admin").Subrouter()
	r.PathPrefix("/api/admin").Handler(negroni.New(
		negroni.HandlerFunc(jwtMiddlewareCommon.HandlerWithNextNoError),
		negroni.HandlerFunc(HandleAdmin),
		negroni.Wrap(admin),
	))
	admin.HandleFunc("/session", getAdminSessions).Methods("GET")
	admin.HandleFunc("/session/{id}", getAdminSession).Methods("GET")
	admin.HandleFunc("/comment/row/{id}", getAdminCommentByRow).Methods("GET")
	admin.HandleFunc("/comment", postComment).Methods("POST")
	admin.HandleFunc("/rate", postRate).Methods("POST")
	admin.HandleFunc("/rate/row/{id}", getAdminRateByRow).Methods("GET")
	admin.HandleFunc("/rate/{id}", postRate).Methods("PUT")
	admin.HandleFunc("/rate/row/{id}/user/me", getAdminRateByRowIdAndUserMe).Methods("GET")
	admin.HandleFunc("/user/notif", postNotifToken).Methods("POST")

	api := mux.NewRouter().PathPrefix("/api/").Subrouter()
	r.PathPrefix("/api").Handler(negroni.New(
		negroni.HandlerFunc(jwtMiddlewareCommon.HandlerWithNextNoError),
		negroni.Wrap(api),
	))
	api.HandleFunc("/application", getApplication).Methods("GET")
	api.HandleFunc("/commonAdmin/currentUser", getCurrentUser).Methods("GET")
	api.HandleFunc("/commonAdmin/login", getAdminUserLogin).Methods("POST")
	api.HandleFunc("/commonAdmin/logout", getAdminUserLogout).Methods("POST")

	// Send all incoming requests to mux.DefaultRouter.
	http.Handle("/", r)
}

func getApplication(w http.ResponseWriter, r *http.Request) {
	application := &Application{
		EventName: config.Get.EVENT_NAME,
		Community: config.Get.COMMUNITY,
		Date: config.Get.DATE,
		ReleaseDate: config.Get.RELEASE_DATE,
		Configured: true,
    }
    json.NewEncoder(w).Encode(application)
	return
}

/**
 * send email
 * @param {[type]} r     *http.Request     [description]
 * @param {[type]} model string            [description]
 * @param {[type]} data  map[string]string data to fo to template
 * @param {[type]} sub   string            subject
 * @param {[type]} to    string            receiver email
 * @param {[type]} from  string            sender email
 */
func SendEmail(r *http.Request, model string, data map[string]string, sub string, to string, from string) error {
		templateFile, err := ioutil.ReadFile("email_templates/" + model)
		if err != nil { return err }
		templateString := string(templateFile)
		tmpl, err := template.New("email").Parse(templateString)
		if err != nil { return err }
		var doc bytes.Buffer 
		err = tmpl.Execute(&doc, data)
		if err != nil { return err }
        body := doc.String()
        c := appengine.NewContext(r)
        msg := &mail.Message{
                Sender:  "Call For Paper <" + from + ">",
                To:      []string{to},
                Subject: sub,
                HTMLBody: body,
        }
        if err := mail.Send(c, msg); err != nil {
            return err
        }
        return nil
}