package cfp

import (
    "encoding/json"
    "os"
)

type configurationProperties struct {
	GOOGLE_SECRET       string
	GITHUB_SECRET       string
	EMAIL_SENDER        string
	JWT_SECRET          string
	NOTIF_SERVER_KEY	string
	EVENT_NAME			string
	COMMUNITY			string
	DATE				string
	RELEASE_DATE		string
	HOSTNAME			string
}

type Config struct {
	Get 				configurationProperties
}

func (v *Config) New() error {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration := configurationProperties{}
	err := decoder.Decode(&configuration)
	if err != nil {
	  return err
	}
	v.Get = configuration
	return nil
}

