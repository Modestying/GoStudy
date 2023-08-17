package http_server

import (
	"fmt"
	"io"
	"net/http"
	_ "net/url"
	"strconv"
)

func HttpServerListen(port int) {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("dd")
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/d", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	fmt.Fprintf(w, "Hello, %s!", r.FormValue("ping"))
	_, err := w.Write([]byte("hello"))
	if err != nil {
		panic("dd")
		return
	}

	fmt.Println(r.FormValue("ping"))

	req, err := http.NewRequest(r.Method, "http://127.0.0.1:81"+r.RequestURI, r.Body)
	//req.Header.Add("Authorization", r.Header.Get("Authorization"))
	client := &http.Client{}
	do, err := client.Do(req)
	if err != nil {
		fmt.Println(do)
		fmt.Println(err.Error())
		return
	}

}
