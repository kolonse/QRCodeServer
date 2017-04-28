// QRCodeServer project main.go
package main

import (
	"flag"
	"net/http"

	"github.com/kolonse/KolonseWeb"
	"github.com/kolonse/KolonseWeb/HttpLib"
	"github.com/kolonse/KolonseWeb/Type"
	qdiy "github.com/kolonse/go-qrcode-extend"
)

var port = flag.Int("p", 18001, "-p=<port> default=18001")
var ip = flag.String("i", "0.0.0.0", "-i=<ip> default=0.0.0.0")

func QRCode(req *HttpLib.Request, res *HttpLib.Response, next Type.Next) {
	defer req.Body.Close()
	var arg qdiy.QRArg
	arg.Parse(req.URL.Query())
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

func Help(req *HttpLib.Request, res *HttpLib.Response, next Type.Next) {
	defer req.Body.Close()
	res.Write([]byte("usage:url?content=<string>&size=<int>&bgcolor=<#int>&forecolor=<#int>&logo=<url>"))
}

func main() {
	flag.Parse()
	KolonseWeb.DefaultApp.Get("/qrcode", QRCode)
	KolonseWeb.DefaultApp.Get("/", Help)
	KolonseWeb.DefaultApp.Post("/qrcode", QRCode)
	KolonseWeb.DefaultApp.Post("/", Help)
	KolonseWeb.DefaultApp.Listen(*ip, *port)
}
