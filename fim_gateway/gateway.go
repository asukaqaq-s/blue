package main

import (
	"encoding/json"
	"fim_server/common/etcd"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"
)

type BaseResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func FilResponse(msg string, res http.ResponseWriter) {
	response := BaseResponse{Code: 7, Msg: msg}
	byteData, _ := json.Marshal(response)
	res.Write(byteData)
}

func auth(authAddr string, res http.ResponseWriter, req *http.Request) (ok bool) {
	authReq, _ := http.NewRequest("POST", authAddr, nil)
	authReq.Header = req.Header
	token := req.URL.Query().Get("token")
	if token != "" {
		authReq.Header.Set("Token", token)
	}
	authReq.Header.Set("ValidPath", req.URL.Path)
	authRes, err := http.DefaultClient.Do(authReq)
	if err != nil {
		logx.Error(err)
		FilResponse("认证服务错误", res)
		return
	}

	type Response struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data *struct {
			UserID uint `json:"userID"`
			Role   int  `json:"role"`
		} `json:"data"`
	}
	var authResponse Response
	byteData, _ := io.ReadAll(authRes.Body)
	authErr := json.Unmarshal(byteData, &authResponse)
	if authErr != nil {
		logx.Error(authErr)
		FilResponse("认证服务错误", res)
		return
	}

	// 认证不通过
	if authResponse.Code != 0 {
		res.Write(byteData)
		return
	}
	if authResponse.Data != nil {
		req.Header.Set("User-ID", fmt.Sprintf("%d", authResponse.Data.UserID))
		req.Header.Set("Role", fmt.Sprintf("%d", authResponse.Data.Role))
	}
	return true
}

type Proxy struct {
}

func (Proxy) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	// 匹配请求前缀  /api/user/xx
	regex, _ := regexp.Compile(`/api/(.*?)/`)
	addrList := regex.FindStringSubmatch(req.URL.Path)
	if len(addrList) != 2 {
		res.Write([]byte("err"))
		return
	}
	service := addrList[1]

	addr := etcd.GetServiceAddr(config.Etcd, service+"_api")
	if addr == "" {
		logx.Errorf("%s 不匹配的服务", service)
		FilResponse("err", res)
		return
	}

	remoteAddr := strings.Split(req.RemoteAddr, ":")
	// 请求认证服务地址
	authAddr := etcd.GetServiceAddr(config.Etcd, "auth_api")
	authUrl := fmt.Sprintf("http://%s/api/auth/authentication", authAddr)
	proxyUrl := fmt.Sprintf("http://%s%s", addr, req.URL.String())

	// 打印日志
	logx.Infof("%s %s", remoteAddr[0], proxyUrl)

	if !auth(authUrl, res, req) {
		return
	}

	remote, _ := url.Parse(fmt.Sprintf("http://%s", addr))
	reverseProxy := httputil.NewSingleHostReverseProxy(remote)
	reverseProxy.ServeHTTP(res, req)
}

var configFile = flag.String("f", "settings.yaml", "the config file")

type Config struct {
	Addr string
	Etcd string
	Log  logx.LogConf
}

var config Config

func main() {
	flag.Parse()
	conf.MustLoad(*configFile, &config)
	logx.SetUp(config.Log)
	fmt.Printf("gateway running %s\n", config.Addr)
	proxy := Proxy{}
	// 绑定服务
	http.ListenAndServe(config.Addr, proxy)
}
