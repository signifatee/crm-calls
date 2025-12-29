package entity

// ARI Channels Post Response Body

type InitCallResponse struct {
	ChannelId    string       `json:"id" db:"channel_id"`
	Name         string       `json:"name" db:"name"`
	State        string       `json:"state" db:"state"`
	ProtocolId   string       `json:"protocol_id" db:"protocol_id"`
	Caller       *CallerIC    `json:"caller"`
	Connected    *ConnectedIC `json:"connected"`
	AccountCode  string       `json:"accountcode" db:"accountcode"`
	Dialplan     *DialplanIC  `json:"dialplan"`
	CreationTime string       `json:"creationtime" db:"creationtime"`
	Language     string       `json:"language" db:"language"`
}

type CallerIC struct {
	Name   string `json:"name" db:"caller_name"`
	Number string `json:"number" db:"caller_number"`
}

type ConnectedIC struct {
	Name   string `json:"name" db:"connected_name"`
	Number string `json:"number" db:"connected_number"`
}

type DialplanIC struct {
	Context  string  `json:"context" db:"dialplan_context"`
	Exten    string  `json:"exten" db:"dialplan_exten"`
	Priority float64 `json:"priority" db:"dialplan_priority"`
	AppName  string  `json:"app_name" db:"dialplan_app_name"`
	AppData  string  `json:"app_data" db:"dialplan_app_data"`
}

func NewICR() *InitCallResponse {
	return &InitCallResponse{
		"",
		"",
		"",
		"",
		&CallerIC{"", ""},
		&ConnectedIC{"", ""},
		"",
		&DialplanIC{"", "", 0, "", ""},
		"",
		"",
	}
}

type InitCallResponseDB struct {
	ChannelId        string  `db:"channel_id"`
	Name             string  `db:"name"`
	State            string  `db:"state"`
	ProtocolId       string  `db:"protocol_id"`
	CallerName       string  `db:"caller_name"`
	CallerNumber     string  `db:"caller_number"`
	ConnectedName    string  `db:"connected_name"`
	ConnectedNumber  string  `db:"connected_number"`
	AccountCode      string  `db:"accountcode"`
	DialplanContext  string  `db:"dialplan_context"`
	DialplanExten    string  `db:"dialplan_exten"`
	DialplanPriority float64 `db:"dialplan_priority"`
	DialplanAppName  string  `db:"dialplan_app_name"`
	DialplanAppData  string  `db:"dialplan_app_data"`
	CreationTime     string  `db:"creationtime"`
	Language         string  `db:"language"`
}

func (icr *InitCallResponse) IcrToICRDB() *InitCallResponseDB {
	return &InitCallResponseDB{
		ChannelId:        icr.ChannelId,
		Name:             icr.Name,
		State:            icr.State,
		ProtocolId:       icr.ProtocolId,
		CallerName:       icr.Caller.Name,
		CallerNumber:     icr.Caller.Number,
		ConnectedName:    icr.Connected.Name,
		ConnectedNumber:  icr.Connected.Number,
		AccountCode:      icr.AccountCode,
		DialplanContext:  icr.Dialplan.Context,
		DialplanExten:    icr.Dialplan.Exten,
		DialplanPriority: icr.Dialplan.Priority,
		DialplanAppName:  icr.Dialplan.AppName,
		DialplanAppData:  icr.Dialplan.AppData,
		CreationTime:     icr.CreationTime,
		Language:         icr.Language,
	}
}

func (icrdb *InitCallResponseDB) IcrDBToICR() *InitCallResponse {
	return &InitCallResponse{
		ChannelId:  icrdb.ChannelId,
		Name:       icrdb.Name,
		State:      icrdb.State,
		ProtocolId: icrdb.ProtocolId,
		Caller: &CallerIC{
			Name:   icrdb.CallerName,
			Number: icrdb.CallerNumber,
		},
		Connected: &ConnectedIC{
			Name:   icrdb.ConnectedName,
			Number: icrdb.ConnectedNumber,
		},
		AccountCode: icrdb.AccountCode,
		Dialplan: &DialplanIC{
			Context:  icrdb.DialplanContext,
			Exten:    icrdb.DialplanExten,
			Priority: icrdb.DialplanPriority,
			AppName:  icrdb.DialplanAppName,
			AppData:  icrdb.DialplanAppData,
		},
		CreationTime: icrdb.CreationTime,
		Language:     icrdb.Language,
	}
}
