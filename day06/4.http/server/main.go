package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	// str := "hello http"
	b, err := ioutil.ReadFile("./xxx.txt")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	// fmt.Println(str)
	w.Write(b)
}

func main() {

	http.HandleFunc("/test", f1)
	http.ListenAndServe("127.0.0.1:9090", nil)

}
