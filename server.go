package main

import (
	"github.com/urfave/negroni"
	"net/http"
	"fmt"
	"html/template"
	""
)

//noinspection ALL
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controller())

	mux.HandleFunc("/aa", func(w http.ResponseWriter, req *http.Request) {
		open(w)
		//fmt.Fprintf(w, "Welcome to the aa page!")
	})
	mux.HandleFunc("/bb", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the bb page!")
	})
	//以上:实例serveMux对象并在此对象中注册handle

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")
}

type HandlerFunc func(rw http.ResponseWriter, r *http.Request)

func controller() HandlerFunc {
	return HandlerFunc(func(a http.ResponseWriter, b *http.Request){
		fmt.Fprintf(a, "zzzzz!")
	})
}

func open(w http.ResponseWriter){
	t,_:= template.ParseFiles("aa.html")
	t.Execute(w,nil)
	fmt.Fprintf(w, "Welcome to the aa page!")
}

//func (n *HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//
//}
