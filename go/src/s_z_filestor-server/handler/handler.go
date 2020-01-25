package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"s_z_filestor-server/meta"
	"s_z_filestor-server/utils"
	"time"
)

func Uploadhandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		fmt.Println(123)
		//bytes, e := ioutil.ReadFile("s_z_filestor-server/static/view/index.html")
		bytes, e := ioutil.ReadFile("s_z_filestor-server/static/view/indexraw.html")
		if e != nil {
			io.WriteString(w, e.Error())
			panic(e)
		}
		io.WriteString(w, string(bytes))

		//w.Write(bytes)
	} else if r.Method == "POST" {

		log.Print("fff")

		file, header, e := r.FormFile("file")

		if e != nil {
			panic("file wrong")
		}

		defer file.Close()

		create, err := os.Create("/tmp/" + header.Filename)
		if err != nil {
			panic("create file wrong")
		}

		defer create.Close()

		meta := Meta.Filemeta{}

		create.Seek(0, 0)

		meta.Filesha1 = utils.FileSha1(create)
		meta.Filename = header.Filename
		meta.Location = "/tmp/" + header.Filename
		meta.UploadAt = time.Now().Format("2006-01-01 12:12:12")
		meta.UploadAt = time.Now().Format("2006-01-02 15:04:05")

		//create.Write()

		meta.Filesize, e = io.Copy(create, file)


		if e != nil {
			panic("write file wrong")
		}

		log.Println("write ok ")
		r.ParseForm()

		fmt.Println(r.Form)

		Meta.Mod(meta)


		fmt.Printf("%+v",meta)
		//UploadSuccHandler(w,r)
		http.Redirect(w, r, "/fileupload/succ", http.StatusFound)
	}

}

func UploadSuccHandler(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "上传ok")
}



func GetimgHandler(w http.ResponseWriter,r *http.Request){

	r.ParseForm()


	fileid:=r.Form["filehash"][0]
	fmeta := Meta.Get(fileid)

	bytes, e := json.Marshal(fmeta)
	if e!=nil{
		w.WriteHeader(http.StatusInternalServerError)

		return
	}


	fmt.Println(r.Form)

	isDownload := r.Form["down"]

	del :=r.Form.Get("del")

	fmt.Println(r.Form)
	if del=="1"{
		delete(Meta.Filemetas,fmeta.Filesha1)

		os.Remove(fmeta.Location)


		w.WriteHeader(http.StatusOK)
		return
	}


	if len(isDownload) >0{
		file,err := os.Open(fmeta.Location)
		if err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))

			return
		}
		defer file.Close()

		all, _ := ioutil.ReadAll(file)

//响应头  当成下载

w.Header().Set("Content-Type","application/octect-stream")
//w.Header().Set("Content-Description","attachment;filename="+fmeta.Filename)
w.Header().Set("content-disposition","attachment;filename=\""+fmeta.Filename+"\"")
		w.Write(all)

	}else{
		w.Write(bytes)

	}







	}

