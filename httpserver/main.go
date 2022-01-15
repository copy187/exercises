package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Printf("http.ListenAndServe err: %v\n", err)
	}
}

func index(resp http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		resp.Header().Set(k, strings.Join(v, " "))
	}
	resp.Header().Set("VERSION", os.Getenv("VERSION"))
	ip := clientIP(req)
	fields := []string{
		time.Now().Format("2006-01-02 15:04:05"),
		ip,
		req.URL.Path,
		req.URL.RawQuery,
	}
	glog.V(1)
	glog.V(2).Info(strings.Join(fields, " "))
	_, _ = resp.Write([]byte("客户端IP：" + ip + "\n" + strings.Join(fields, " ") + "\n"))
	resp.WriteHeader(http.StatusOK)
}
func clientIP(r *http.Request) (ip string) {
	ip = strings.TrimSpace(r.Header.Get("X-Forwarded-For"))
	if ip = strings.Split(ip, ",")[0]; ip != "" {
		return
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return
	}
	ip = r.RemoteAddr
	ip, _, _ = net.SplitHostPort(ip)
	return
}
func healthz(resp http.ResponseWriter, req *http.Request) {
	j,_ :=json.Marshal(map[string]string{"a":"𥖄"})
	resp.Write(j)
}
