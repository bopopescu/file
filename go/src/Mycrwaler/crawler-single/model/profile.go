package model

import "encoding/json"

type Profile struct {
	Name  string
	 Age, Marriage  string
	Sex,Img                                   string
	Income                                string
}
// todo Map 转 Struce
// Deprecated
func Map2Profile(o interface{}) Profile  {
	str, _ := json.Marshal(o)
	p := Profile{}
	json.Unmarshal(str, &p)
	return p
}