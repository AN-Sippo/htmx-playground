package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter,r *http.Request){
	reader,err := os.Open("index.html");
	if err != nil {
		panic(err);
	}
	io.Copy(w,reader);
}

func buttonHandler(w http.ResponseWriter,r *http.Request){
	reader,err := os.Open("buttonClicked.html");
	if err != nil {
		panic(err);
	}
	io.Copy(w,reader);
}

func main(){
	server := http.Server{
		Addr: ":8080",
		Handler: nil,
	}
	fmt.Println("listening at :8080...")
	http.HandleFunc("/",handler)
	http.HandleFunc("/button",buttonHandler)
	server.ListenAndServe();
}
