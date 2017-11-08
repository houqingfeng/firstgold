package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

//server struct
type LeVSHttpServer struct {
	server       *http.Server
	serverPath   string
}

func (httpServer *LeVSHttpServer) startTransformServer(sourcePath string) {
	log.Print("start server >> " + sourcePath)
	if strings.HasSuffix(sourcePath, "/") == false {
		sourcePath += "/"
	}

	httpServer.serverPath = sourcePath
	mux := http.NewServeMux()
	httpServer.server = &http.Server{
		Addr:           ":8128",
		Handler:        mux,
		ReadTimeout:    3600 * time.Second * 1 , //传输最长时间一个小时
		WriteTimeout:   3600 * time.Second * 1 , //传输最长时间一个小时
		MaxHeaderBytes: 1000,
	}

	//http.HandleFunc("/userInfo/", httpServer.userInfoServer)   // 用户信息服务器

	mux.HandleFunc("/videoInfo/", httpServer.videoInfoServer) // 文件信息服务器
	mux.HandleFunc("/test/", httpServer.upload)
	mux.HandleFunc("/upload/", httpServer.upload)
	// 请求文件
	mux.HandleFunc("/", httpServer.uploadHandle)                  // 请求文件
	mux.HandleFunc("/dir/", httpServer.upLoadDirectory)           // 请求文件夹
	mux.HandleFunc("/dir/files/", httpServer.uploadDirectorFiles) // 请求文件夹文件
	mux.HandleFunc("/cancel/", httpServer.cancel)

	mux.HandleFunc("/videostream/", httpServer.videoStream)

	err := httpServer.server.ListenAndServe()
	if err != nil {
		log.Print(err)
	}
}

var test = true
func (httpServer *LeVSHttpServer) videoStream(w http.ResponseWriter, r *http.Request) {
	log.Print("请求的ip地址为：", r.RemoteAddr)

	if test {
		file, _ := ioutil.ReadFile("/home/hou/Downloads/7.mp4")
		w.Write(file)
	} else {
		video, err := os.Open("/home/hou/Downloads/7.mp4")
		if err != nil{
			log.Println("Open file failed!")
			w.Write([]byte("Failed"))
		}

		buf := make([]byte, 32*1024)
		for {
			size, errRead := video.Read(buf)
			w.Write(buf[0:size])
			if errRead != nil {
				break
			}
		}
	}
}

func (httpServer *LeVSHttpServer) cancel(w http.ResponseWriter, r *http.Request) {

}

func (httpServer *LeVSHttpServer) upload(w http.ResponseWriter, r *http.Request) {
	log.Print("请求的ip地址为：", r.RemoteAddr)
	body, _ := ioutil.ReadAll(r.Body)
	str := string(body[:])
	log.Print("接收到的str: ", str)
	w.Header()
	w.Write([]byte("hello welcome"))
}


func (httpServer *LeVSHttpServer) videoInfoServer(w http.ResponseWriter, r *http.Request) {

}

func (httpServer *LeVSHttpServer) uploadHandle(w http.ResponseWriter, r *http.Request) {
}


func (httpServer *LeVSHttpServer) uploadDirectorFiles(w http.ResponseWriter, r *http.Request) {

}

func (httpServer *LeVSHttpServer) upLoadDirectory(w http.ResponseWriter, r *http.Request) {

}

func (httpServer *LeVSHttpServer) stopTransformServer() {

}

func (httpServer *LeVSHttpServer) StartHttpServerAsync() {

}

func (httpServer *LeVSHttpServer) StopHttpServer() {

}

func main()  {
	server := &LeVSHttpServer{}
	server.startTransformServer("/")
}