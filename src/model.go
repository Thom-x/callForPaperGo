package cfp

type Session struct {
	ID int64 `datastore:"-" json:"id"`
    Email string `json:"email" validate:"nonzero"`
    Name string `json:"name" validate:"nonzero"`
    Firstname string `json:"firstname" validate:"nonzero"`
    Phone string `json:"phone"`
    Company string `json:"company"`
    Bio string `json:"bio" validate:"nonzero"`
    Social string `json:"social"`
    Type string `json:"type" validate:"nonzero"`
    SessionName string `json:"sessionName" validate:"nonzero"`
    Description string `json:"description" validate:"nonzero"`
    References string `json:"references"`
    Difficulty int `json:"difficulty" validate:"min=1,max=3,nonzero"`
    Track string `json:"track" validate:"nonzero"`
    CoSpeaker string `json:"coSpeaker"`
    Financial bool `json:"financial"`
    Travel bool `json:"travel"`
    TravelFrom string `json:"travelFrom"`
    Hotel bool `json:"hotel"`
    HotelFrom string `json:"hotelFrom"`
    Draft bool `json:"draft"`
    Added int64 `json:"added"`
}

type Draft struct {
	ID int64 `datastore:"-" json:"id"`
    Email string `json:"email"`
    Name string `json:"name"`
    Firstname string `json:"firstname"`
    Phone string `json:"phone"`
    Company string `json:"company"`
    Bio string `json:"bio"`
    Social string `json:"social"`
    Type string `json:"type"`
    SessionName string `json:"sessionName"`
    Description string `json:"description"`
    References string `json:"references"`
    Difficulty int `json:"difficulty" validate:"min=0,max=3"`
    Track string `json:"track"`
    CoSpeaker string `json:"coSpeaker"`
    Financial bool `json:"financial"`
    Travel bool `json:"travel"`
    TravelFrom string `json:"travelFrom"`
    Hotel bool `json:"hotel"`
    HotelFrom string `json:"hotelFrom"`
    Draft bool `json:"draft"`
    Added int64 `json:"added"`
}

type AdminComment struct {
	ID int64 `datastore:"-" json:"id"`
	RowId string `json:"rowId" validate:"nonzero"`
	Added int64 `json:"added"`
    Comment string `json:"comment" validate:"nonzero"`
}

type AdminCommentResponse struct {
	ID int64 `datastore:"-" json:"id"`
	Added int64 `json:"added"`
	RowId string `json:"rowId"`
	User AdminUser `json:"user"`
    Comment string `json:"comment"`
}

type AdminRate struct {
	ID int64 `datastore:"-" json:"id"`
	RowId string `json:"rowId" validate:"nonzero"`
    Rate int `json:"rate" validate:"min=1,max=5,nonzero"`
}

type AdminRateReponse struct {
	ID int64 `datastore:"-" json:"id"`
	RowId string `json:"rowId"`
	User AdminUser `json:"user"`
    Rate int `json:"rate" validate:"min=1,max=5,nonzero"`
}

type Profile struct {
    Name string `json:"name"`
    Firstname string `json:"firstname"`
    Email string `json:"email"`
    Phone string `json:"phone"`
    Company string `json:"company"`
    Bio string `json:"bio"`
    Social string `json:"social"`
}

type LoginUser struct {
    Email string `json:"email" validate:"nonzero"`
    Password string `json:"password" validate:"nonzero"`
}

type SignUpUser struct {
    Email string `json:"email" validate:"nonzero"`
    Password string `json:"password" validate:"nonzero"`
    Captcha string `json:"captcha" validate:"nonzero"`
}

type Token struct {
    Token string `json:"token"`
}

type User struct {
	ID int64 `datastore:"-" json:"id"`
    Email string `json:"email"`
    Password string `json:"password"`
    Profile string `json:"profile"`
    Google string `json:"google"`
    Github string `json:"github"`
    Verified bool `json:"verified"`
    TokenVerified string `json:"tokenVerified"`
}

type AdminUser struct {
	ID int64 `datastore:"-" json:"id"`
    Email string `json:"email"`
    Name string `json:"name"`
    Notif string `json:"notif"`
}

type Application struct {
	EventName string `json:"eventName"`
	Community string `json:"community"`
	Date string `json:"date"`
	ReleaseDate string `json:"releaseDate"`
	Configured bool `json:"configured"`
}

type CurrentUser struct {
	Connected bool `json:"connected"`
	Admin bool `json:"admin"`
	Config bool `json:"config"`
	Uri string `json:"uri"`
}

type Redirect struct {
	Redirect string `json:"redirect"`
}

type Uri struct {
	Uri string `json:"uri"`
}


type OAuth2Params struct {
	Code         string `json:"code" url:"code"`
	ClientId     string `json:"client_id" url:"client_id"`
	ClientSecret string `json:"client_secret" url:"client_secret"`
	RedirectUri  string `json:"redirect_uri" url:"redirect_uri"`
	GrantType    string `json:"grant_type,omitempty" url:"grant_type,omitempty"`
}

type accessTokenData struct {
	AccessToken string `json:"access_token" url:"access_token"`
	TokenType   string `json:"token_type" url:"token_type"`
	ExpiresIn   int    `json:"expires_in" url:"expires_in"`
}

type Notif struct{
    Ids []string `json:"registration_ids" url:"registration_ids"`
}
