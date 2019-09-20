package message

const (
	Loginmestype ="LoginMes"
	Loginmesrestype ="LoginMesRes"
	RegMesType ="RegMes"
)
type Message struct {
	Type string `json:"type"`
	Data string	`json:"data"`
}




type LoginMes struct {
	Userid int	`json:"userid"`
	Userpwd string `json:"userwd"`
	Username string `json:"usernm"`
}

type LoginMesRes struct {
	Code int `json:"code"`
	Error string `json:"msg"`

}


type RegMes struct {
	Code int `json:"code"`
	Error string `json:"msg"`

}