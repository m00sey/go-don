package api

import "net/http"

func greetingHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(200)
	_, _ = res.Write([]byte("hi"))
}