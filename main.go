// QRCodeServer project main.go
package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"flag"
	"net/http"

	"github.com/kolonse/KolonseWeb"
	"github.com/kolonse/KolonseWeb/HttpLib"
	"github.com/kolonse/KolonseWeb/Type"
	qdiy "github.com/kolonse/go-qrcode-extend"
	"github.com/kolonse/simplekv"
)

var port = flag.Int("p", 18001, "-p=<port> default=18001")
var ip = flag.String("i", "0.0.0.0", "-i=<ip> default=0.0.0.0")
var storeDir = flag.String("d", "/tmp/simplekv", "-d=<dir> 活码内容存储目录")
var storeName = flag.String("n", "qrcode", "-n=<name> 活码存储目录名字")
var activeUrl = flag.String("a", "http://qrcode.youdianle.com.cn/active", "-a=<url> 连接性活码地址url")

type Shortest struct {
	Redirect bool
	Content  string
}

var kv *simplekv.SKV

func QRCode(req *HttpLib.Request, res *HttpLib.Response, next Type.Next) {
	defer req.Body.Close()
	s := req.URL.Query().Get("shortest")
	query := req.URL.Query()
	if s == "true" { // 是否使用活码
		s = query.Get("redirect")
		st := Shortest{
			Redirect: false,
			Content:  query.Get("content"),
		}
		if s == "true" { // 是否对内容重定向
			st.Redirect = true
		}
		ctr := md5.New()
		ctr.Write([]byte(st.Content))
		rs := hex.EncodeToString(ctr.Sum(nil))
		if !kv.Exist(rs) {
			jbuf, _ := json.Marshal(st)
			kv.Write(rs, jbuf)
		}
		query.Set("content", *activeUrl+"?k="+rs)
	}

	var arg qdiy.QRArg
	arg.Parse(query)
	if len(arg.Content) == 0 {
		http.NotFound(res.ResponseWriter, req.Request)
		KolonseWeb.BeeLogger.Error("%s %s %s", req.RemoteAddr, req.URL.String(), "不支持空内容二维码请求")
		return
	}

	qlogo := qdiy.QRDiy{Arg: arg}
	buf, err := qlogo.Encode()
	if err != nil {
		http.NotFound(res.ResponseWriter, req.Request)
		KolonseWeb.BeeLogger.Error("%s %s %s", req.RemoteAddr, req.URL.String(), err.Error())
		return
	}
	res.Write(buf)
	KolonseWeb.BeeLogger.Info("%s %s %s", req.RemoteAddr, req.URL.String(), "success")
}

func Active(req *HttpLib.Request, res *HttpLib.Response, next Type.Next) {
	defer req.Body.Close()
	key := req.URL.Query().Get("k")
	if !kv.Exist(key) {
		http.NotFound(res.ResponseWriter, req.Request)
		KolonseWeb.BeeLogger.Error("%s %s %s NotFound", req.RemoteAddr, req.URL.String(), key)
		return
	}
	obj := kv.Read(key).ToJsonObject()
	st := Shortest{
		Redirect: obj["Redirect"].(bool),
		Content:  obj["Content"].(string),
	}
	KolonseWeb.BeeLogger.Info(key)
	if st.Redirect { // 如果属于重定向 那么需要将请求内容返回给请求源hex
		res.Redirect(string(st.Content))
	} else {
		res.End(st.Content)
	}
	KolonseWeb.BeeLogger.Info("%s %s %s", req.RemoteAddr, req.URL.String(), "success")
}

func Help(req *HttpLib.Request, res *HttpLib.Response, next Type.Next) {
	defer req.Body.Close()
	res.Write([]byte("usage:url?content=<string>&size=<int>&bgcolor=<#int>&forecolor=<#int>&logo=<url>"))
}

func main() {
	var err error
	kv, err = simplekv.NewSKV(*storeName, *storeDir)
	if err != nil {
		panic(err)
	}
	flag.Parse()
	KolonseWeb.DefaultApp.Get("/qrcode", QRCode)
	KolonseWeb.DefaultApp.Get("/active", Active)
	KolonseWeb.DefaultApp.Get("/", Help)
	KolonseWeb.DefaultApp.Post("/qrcode", QRCode)
	KolonseWeb.DefaultApp.Post("/active", Active)
	KolonseWeb.DefaultApp.Post("/", Help)
	KolonseWeb.DefaultApp.Listen(*ip, *port)
}
