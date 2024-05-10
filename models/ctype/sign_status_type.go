package ctype

import "encoding/json"

type SignStatus int

const (
	SignQQ    SignStatus = 1 // qq
	SignGitee SignStatus = 2 // gitee
	SignEmail SignStatus = 3 // email
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	var str string
	switch s {
	case SignQQ:
		str = "QQ"
	case SignGitee:
		str = "Gitee"
	case SignEmail:
		str = "Email"
	default:
		str = "其它"
	}
	return str
}
