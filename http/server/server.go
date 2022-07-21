package main

import (
	"apis/http/logger"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func callLogger(obj logger.LogInstance) {
	if obj == nil {
		return
	}
	logger.Logger("server", obj)
}

func getData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "your'e in getData api %s and method %s\n", "of http", r.Method)
}

func postData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "your'e in postData api %s and method %s\n", "of http", r.Method)
	//prefer io.Copy over ioutil.ReadAll to avoid memory spikes
	n, err := io.Copy(w, r.Body)
	callLogger(err)
	fmt.Fprintf(w, "\n%d number of bytes are written\n", n)
	r.Body.Close()
}

func postMultipartData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "your'e in postMultipart api %s and method %s\n", "of http", r.Method)

	//ParseForm populates Form and PostFrom
	r.ParseForm()
	fmt.Println(r.PostForm)
	val1 := r.Form
	for k, v := range val1 {
		fmt.Fprintf(w, "%s : %s\n", k, v)
		fmt.Println(k, v)
	}
	val2 := r.PostForm
	for k, v := range val2 {
		fmt.Fprintf(w, "%s : %s\n", k, v)
		fmt.Println(k, v)
	}

	n, err := io.Copy(w, r.Body)
	callLogger(err)
	fmt.Fprintf(w, "\n%d number of bytes are written\n", n)
	r.Body.Close()
}

func main() {

	callLogger(errors.New("In server"))

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		getData(w, r)
	})

	http.HandleFunc("/post/", func(w http.ResponseWriter, r *http.Request) {
		postData(w, r)
	})

	http.HandleFunc("/postForm/", func(w http.ResponseWriter, r *http.Request) {
		postMultipartData(w, r)
	})

	fs := http.FileServer(http.Dir("."))

	//Here if StripPrefix isn't used, then  let the url be /files/ex1.txt and it checks for /files/ex1.txt relative to the mentioned ./files/ex1.txt
	http.Handle("/files/", http.StripPrefix("/files/", fs))

	err := http.ListenAndServe(":3333", nil)
	callLogger(err)
}
