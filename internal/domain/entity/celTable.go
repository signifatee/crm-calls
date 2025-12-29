package entity

type CelTable struct {
	Id          int    `json:"id" db:"id"`
	Eventtype   string `json:"eventtype" db:"eventtype"`
	Eventtime   string `json:"eventtime" db:"eventtime"`
	Cid_name    string `json:"cid_name,omitempty" db:"cid_name"`
	Cid_num     string `json:"cid_num,omitempty" db:"cid_num"`
	Cid_ani     string `json:"cid_ani,omitempty" db:"cid_ani"`
	Cid_rdnis   string `json:"cid_rdnis,omitempty" db:"cid_rdnis"`
	Cid_dnid    string `json:"cid_dnid,omitempty" db:"cid_dnid"`
	Exten       string `json:"exten" db:"exten"`
	Context     string `json:"context" db:"context"`
	Channame    string `json:"channame" db:"channame"`
	Appname     string `json:"appname" db:"appname"`
	Appdata     string `json:"appdata" db:"appdata"`
	Amaflags    string `json:"amaflags" db:"amaflags"`
	Accountcode string `json:"accountcode,omitempty" db:"accountcode"`
	Uniqueid    string `json:"uniqueid" db:"uniqueid"`
	Linkedid    string `json:"linkedid" db:"linkedid"`
	Peer        string `json:"peer,omitempty" db:"peer"`
	Userdeftype string `json:"userdeftype,omitempty" db:"userdeftype"`
	Extra       string `json:"extra" db:"extra"`
}
