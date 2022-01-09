package main

import (
	"fmt"
	"github.com/golang/glog"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil{
		fmt.Printf("http.ListenAndServe err: %v\n", err)
	}
}

func index(resp http.ResponseWriter,req *http.Request){
	for k,v := range req.Header{
		resp.Header().Set(k, strings.Join(v, " "))
	}
	resp.Header().Set("VERSION", os.Getenv("VERSION"))

	fields := []string{
		time.Now().Format("2006-01-02 15:04:05"),
		req.RemoteAddr,
		req.URL.Path,
		req.URL.RawQuery,
	}
	glog.V(4)
	glog.V(2).Info(strings.Join( fields, " "))
	_,_ = resp.Write([]byte("客户端IP："+req.RemoteAddr+"\n"))
	_,_ = resp.Write([]byte(strings.Join( fields, " ")+"\n"))

	resp.WriteHeader(http.StatusOK)
}

func healthz(resp http.ResponseWriter,req *http.Request) {
	resp.WriteHeader(http.StatusOK)
}