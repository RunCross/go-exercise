package main

import (
	// "bytes"
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	HOST   = "http://resyy.changyan.com/resourceToolRelease"
	LIST   = HOST + "/index.php?c=resourceModifyAction&a=getResourceList"
	DELOPY = HOST + "/index.php?c=resourceUploadAction&a=publishResource"
)

func main() {
	buffer := "&attributes[press]=&attributes[grade]=1&attributes[district]=&attributes[version]=&attributes[addname]=&attributes[period]=&attributes[restype]=&states[]=1&flag=deploy"
	// body := bytes.NewBuffer([]byte(buffer))
	fmt.Println(LIST)
	fmt.Println(buffer)
	resget, err := http.Get(LIST + "&" + buffer)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(resget)
	result, err := ioutil.ReadAll(resget.Body)
	resget.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s", result)
	/*res, err := http.Post(LIST, "application/json;charset=utf-8", body)
	if err != nil {
		log.Fatal(err)
		return
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s", result)*/
	bio := bufio.NewReader(os.Stdin)
	line, hasMoreInLine, err := bio.ReadLine()
	fmt.Println(string(line), hasMoreInLine)
	i := 1
choice:
	i++
	if i < 2 {
		goto choice
	}
	fmt.Println(i)
}
