package dtmcli

// Msg reliable msg type
type Msg struct {
	TransBase
	Steps         []MsgStep `json:"steps"`
	QueryPrepared string    `json:"query_prepared"`
}

// MsgStep struct of one step msg
type MsgStep struct {
	Action string `json:"action"`
	Data   string `json:"data"`
}

// NewMsg create new msg
func NewMsg(server string, gid string) *Msg {
	return &Msg{
		TransBase: TransBase{
			Gid:       gid,
			TransType: "msg",
			Dtm:       server,
		},
	}
}

// Add add a new step
func (s *Msg) Add(action string, postData interface{}) *Msg {
	Logf("msg %s Add %s %v", s.Gid, action, postData)
	step := MsgStep{
		Action: action,
		Data:   MustMarshalString(postData),
	}
	s.Steps = append(s.Steps, step)
	return s
}

// Prepare prepare the msg
func (s *Msg) Prepare(queryPrepared string) error {
	s.QueryPrepared = OrString(queryPrepared, s.QueryPrepared)
	return s.CallDtm(s, "prepare")
}

// Submit submit the msg
func (s *Msg) Submit() error {
	return s.CallDtm(s, "submit")
}
