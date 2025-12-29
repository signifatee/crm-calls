package entity

type CdrTable struct {
	Calldate      string `json:"eventtype,omitempty"`
	Clid          string `json:"clid,omitempty"`
	Src           string `json:"src,omitempty"`
	Dst           string `json:"dst,omitempty"`
	Dcontext      string `json:"dcontext,omitempty"`
	Channel       string `json:"channel,omitempty"`
	Dstchannel    string `json:"dstchannel,omitempty"`
	Lastapp       string `json:"lastapp,omitempty"`
	Lastdata      string `json:"lastdata,omitempty"`
	Duration      string `json:"duration,omitempty"`
	Billsec       string `json:"billsec,omitempty"`
	Disposition   string `json:"disposition,omitempty"`
	Amaflags      string `json:"amaflags,omitempty"`
	Accountcode   string `json:"accountcode,omitempty"`
	Uniqueid      string `json:"uniqueid,omitempty"`
	Userfield     string `json:"userfield,omitempty"`
	Did           string `json:"did,omitempty"`
	Recordingfile string `json:"recordingfile,omitempty"`
	Cnum          string `json:"cnum,omitempty"`
	Cnam          string `json:"cnam,omitempty"`
	Outbound_cnum string `json:"outbound_cnum,omitempty"`
	Outbound_cnam string `json:"outbound_cnam,omitempty"`
	Dst_cnam      string `json:"dst_cnam,omitempty"`
	Linkedid      string `json:"linkedid,omitempty"`
	Peeraccount   string `json:"peeraccount,omitempty"`
	Sequence      string `json:"sequence,omitempty"`
}
