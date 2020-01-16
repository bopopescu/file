package model


//go 的结构体


type SearchResult struct {
	Hits     int64
	Start    int
	Query    string
	PrevFrom int
	NextFrom int
	Items    []interface{}
}
