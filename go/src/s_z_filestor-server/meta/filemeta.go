package Meta


type Filemeta struct {

	Filesha1 string
	Filename string
	Filesize int64
	Location  string
	UploadAt string
}


var Filemetas map[string]Filemeta

func init(){

	Filemetas=make(map[string]Filemeta )
}


func Mod(filemeta Filemeta){
	Filemetas[filemeta.Filesha1]= filemeta

}

func Get(fileid string) Filemeta{

	return Filemetas[fileid]
}