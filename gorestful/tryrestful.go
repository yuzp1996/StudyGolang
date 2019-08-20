package main

import (
	"github.com/emicklei/go-restful"
	"io"
	"log"
	"net/http"
)

func main(){
	ws := new(restful.WebService)
	ws.Route(ws.GET("/apis/auth.alauda.io/v1/projects").To(hello))
	ws.Route(ws.GET("/apis/auth.alauda.io/v1/projects1").To(hello))

	restful.Add(ws)

	log.Fatal(http.ListenAndServe(":8080",nil))

}

func hello(req *restful.Request, resp *restful.Response){
	_,_=io.WriteString(resp, "world")
}
