package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const proxyAddr string = "localhost:9000"

var (
	counter            = 0
	firstInstanceHost  = "http://localhost:8080"
	secondInstanceHost = "http://localhost:8082"
)

func main() {
	http.HandleFunc("/", handleProxy)
	log.Fatalln(http.ListenAndServe(proxyAddr, nil))
}

func handleProxy(w http.ResponseWriter, r *http.Request) {
	textBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln()
	}
	text := string(textBytes)

	if counter == 0 {
		resp, err := http.Post(firstInstanceHost, "text/plain", bytes.NewBuffer([]byte(text)))
		if err != nil {
			log.Fatalln(err)
		}
		counter++
		textBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		fmt.Println(string(textBytes))
		return
	}
	resp, err := http.Post(secondInstanceHost, "text/plain", bytes.NewBuffer([]byte(text)))
	if err != nil {
		log.Fatalln(err)
	}
	textBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	fmt.Println(string(textBytes))
	counter--
}
