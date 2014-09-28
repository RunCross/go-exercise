package main

import (
	"fmt"
	// "io"
	"html/template"
	"log"
	"net/http"
	"strings"
	// "server/ser"
	// "time"
)

type iHandler struct {
}

func main() {

	// myHandler := &iHandler{}
	// s := &http.Server{
	// 	Addr: ":9393",
	// 	// Handler:        myHandler,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	http.NewServeMux().HandleFunc("/css", func(rw http.ResponseWriter, req *http.Request) {
		file := "./css" + req.URL.Path[len("/css")-1:]
		// if(0 & 0x0001)==0{
		// 	if
		// }
		http.ServeFile(rw, req, file)
	})

	http.HandleFunc("/test", myHandler)
	http.HandleFunc("/nav", navHandler)
	log.Fatal(http.ListenAndServe(":9393", nil))
	// log.Fatal(s.ListenAndServe())
}

func myHandler(rw http.ResponseWriter, req *http.Request) {

	// log.Fatalln(req.URL)
	// 解析参数, 默认是不会解析的
	// log.Fatalln(req.ParseForm())
	fmt.Printf("req.ParseForm(): %v\n", req.ParseForm())
	// 这些信息是输出到服务器端的打印信息
	// log.Fatalln("request map:", req.Form)
	// log.Fatalln("path", req.URL.Path)
	// log.Fatalln("scheme", req.URL.Scheme)
	// log.Fatalln(req.Form["url_long"])
	// log.Fatalln(req.FormValue("go"))
	// fmt.Println(req.FormValue("go"))
	for k, v := range req.Form {
		fmt.Printf("key: %s", k)
		fmt.Printf("val: %s", strings.Join(v, ";"))
	}

	// 这个写入到w的信息是输出到客户端的
	fmt.Fprintf(rw, "Hello ")
}

func (*iHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// rw.Write([]byte("Ihand:Ihand"))
}

func navHandler(rw http.ResponseWriter, req *http.Request) {
	// http.Redirect(rw, req, "/Temaplate/index.html", http.StatusFound)
	t, err := template.ParseFiles("index.html")
	fmt.Println(req.URL)

	if err == nil {
		t.Execute(rw, nil)
	} else {
		fmt.Println(err.Error())
	}
}
