package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	// "flag"
	flag "github.com/spf13/pflag"
	"fmt"
)

// NewProxy 拿到 targetHost 后，创建一个反向代理
func NewProxy(targetHost string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(url)
	return proxy, nil
}

// ProxyRequestHandler 使用 proxy 处理请求
func ProxyRequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}


var port int
var listen string
var target string

func init() {
    flag.IntVarP(&port, "port", "p", 80, "listen on port")
    flag.StringVarP(&listen, "listen", "l", "0.0.0.0", "listening IP")
    flag.StringVarP(&target, "target", "t", "", "target URL")
}

func main() {
	flag.Parse()

	// 初始化反向代理并传入真正后端服务的地址
	proxy, err := NewProxy(target)

	if err != nil {
		panic(err)
	}

	// handle all requests to your server using the proxy
	// 使用 proxy 处理所有请求到你的服务
	http.HandleFunc("/", ProxyRequestHandler(proxy))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d",listen,port), nil))
}
