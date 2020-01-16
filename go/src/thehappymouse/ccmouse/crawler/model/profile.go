package model

import "encoding/json"

type Profile struct {
	Name, Height, Weight, Job, JobAddress string
	Edu, HasChild, Jiguan, Age, Marriage  string
	Sex                                   string
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