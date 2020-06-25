package engine

import (
	sync2 "sync"
)

var du1pmap = make(map[string]string)


var du1pmap2 = make(map[string]string)

var sync sync2.Mutex
var sync3 sync2.Mutex


func IsDuplicate2(url string) bool {

	sync.Lock()
	defer sync.Unlock()
	_,ok:=du1pmap[url]
	//fmt.Println(url,"dup",s,ok)
	if ok{
		return true
	}else{
		du1pmap[url]=url

	}
	return false
}

func IsDuplicate23(url string) bool {

	sync3.Lock()
	defer sync3.Unlock()
	_,ok:=du1pmap2[url]
	//fmt.Println(url,"dup",s,ok)
	if ok{
		return true
	}else{
		du1pmap2[url]=url

	}
	return false
}