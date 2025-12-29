package dto

type ChannelDestroyed struct {
	Type      string `json:"type"`
	Timestamp string `json:"timestamp"`
	Cause     int    `json:"cause"`
	CauseTxt  string `json:"cause_txt"`
	Channel   struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		State  string `json:"state"`
		Caller struct {
			Name   string `json:"name"`
			Number string `json:"number"`
		} `json:"caller"`
		Connected struct {
			Name   string `json:"name"`
			Number string `json:"number"`
		} `json:"connected"`
		Accountcode string `json:"accountcode"`
		Dialplan    struct {
			Context  string `json:"context"`
			Exten    string `json:"exten"`
			Priority int    `json:"priority"`
			AppName  string `json:"app_name"`
			AppData  string `json:"app_data"`
		} `json:"dialplan"`
		Creationtime string `json:"creationtime"`
		Language     string `json:"language"`
	} `json:"channel"`
	AsteriskId  string `json:"asterisk_id"`
	Application string `json:"application"`
}
