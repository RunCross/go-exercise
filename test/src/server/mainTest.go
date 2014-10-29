package main

import (
	"fmt"
	// "io"
	"html/template"
	"log"
	"net/http"
	"os"
	//"strings"
	// "server/ser"
	"time"
)

type iHandler struct {
}

const (
	ListDir = 0x0001
)

func main() {

	// myHandler := &iHandler{}
	// s := &http.Server{
	// 	Addr: ":9393",
	// 	// Handler:        myHandler,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// http.NewServeMux().HandleFunc("/css", func(rw http.ResponseWriter, req *http.Request) {
	// 	file := "./css" + req.URL.Path[len("/css")-1:]
	// 	// if(0 & 0x0001)==0{
	// 	// 	if
	// 	// }
	// 	http.ServeFile(rw, req, file)
	// })
	mux := http.NewServeMux()
	staticDirHandler(mux, "/css/", "./css", 0)
	// mux.HandleFunc("/css/", func(rw http.ResponseWriter, req *http.Request) {
	// 	fmt.Println("css")
	// })

	mux.HandleFunc("/test", myHandler)
	mux.HandleFunc("/nav", navHandler)
	mux.HandleFunc("/exam", examHandler)
	log.Fatal(http.ListenAndServe(":9393", mux))
	// log.Fatal(s.ListenAndServe())
}

func myHandler(rw http.ResponseWriter, req *http.Request) {

	// log.Fatalln(req.URL)
	// 解析参数, 默认是不会解析的
	// log.Fatalln(req.ParseForm())
	// fmt.Printf("req.ParseForm(): %v\n", req.ParseForm())
	// fmt.Println(time.Now().String() + "  " + req.URL.Path)
	// 这些信息是输出到服务器端的打印信息
	// log.Fatalln("request map:", req.Form)
	// log.Fatalln("path", req.URL.Path)
	// log.Fatalln("scheme", req.URL.Scheme)
	// log.Fatalln(req.Form["url_long"])
	// log.Fatalln(req.FormValue("go"))
	// fmt.Println(req.FormValue("go"))
	// for k, v := range req.Form {
	// 	fmt.Printf("key: %s", k)
	// 	fmt.Printf("val: %s", strings.Join(v, ";"))
	// }

	// 这个写入到w的信息是输出到客户端的
	fmt.Fprintf(rw, "Hello ")
}

func (*iHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// rw.Write([]byte("Ihand:Ihand"))
}

func navHandler(rw http.ResponseWriter, req *http.Request) {
	// http.Redirect(rw, req, "/Temaplate/index.html", http.StatusFound)
	t, err := template.ParseFiles("Template/index.html")
	fmt.Println(time.Now().String() + "  " + req.URL.Path)

	if err == nil {
		t.Execute(rw, nil)
	} else {
		fmt.Println(err.Error())
	}
}

func examHandler(rw http.ResponseWriter, req *http.Request) {
	// http.Redirect(rw, req, "/Temaplate/index.html", http.StatusFound)
	t, err := template.ParseFiles("Template/example.html")
	fmt.Println(time.Now().String() + "  " + req.URL.Path)

	if err == nil {
		t.Execute(rw, nil)
	} else {
		fmt.Println(err.Error())
	}
}

func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {
	fmt.Println("mux")
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		fmt.Println(file)
		if (flags & ListDir) == 0 {
			if exists := isExists(file); !exists {
				http.NotFound(w, r)
				return
			}
		}
		http.ServeFile(w, r, file)
	})
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}
