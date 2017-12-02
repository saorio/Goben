package main

import (
	"net/http"
	"time"

	"github.com/elazarl/goproxy"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.OnRequest(goproxy.DstHostIs("www.twitter.com")).DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			if h, _, _ := time.Now().Clock(); h >= 8 && h <= 17 {
				return r, goproxy.NewResponse(r,
					goproxy.ContentTypeText, http.StatusForbidden,
					"Don't waste your time!")
			}
			return r, nil
		})
}
