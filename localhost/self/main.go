package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
	//"net/http"
	"db"
)

func main() {
	//初始化app
	dotapp := dotweb.New()

	dotapp.SetLogPath("/Users/liangsijun/go")

	InitRouter(dotapp.HttpServer)

	db.Insert()
	err := dotapp.StartServer(3000)
	fmt.Println("dotweb.StartServer error => ", err)

}


func Index(ctx *dotweb.HttpContext) {
	ctx.View("/Users/liangsijun/go/localhost/self/template/index.html")
}

func InitRouter(server *dotweb.HttpServer) {
	server.Router().GET("/index", Index)
	server.Router().ServerFile("/static/*filepath",  "static")
}