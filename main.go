package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const address string = ":1111"

func payments(w http.ResponseWriter, r *http.Request) {

	/*
	   b, err := ioutil.ReadAll(r.Body)
	   if err != nil {
	       log.Printf("Error reading body: %v", err)
	       http.Error(w, "can't read body", http.StatusBadRequest)
	       return
	   }
	   defer r.Body.Close()
	   fmt.Println("body: ", b)
	*/
	//fmt.Fprintf(w, "ioutil.ReadAll(r.Body) is %v: \n", body)

	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Println("body is:", string(b))

}

func main() {

	http.HandleFunc("/v1/payments", payments)

	log.Println("[main] listening on port", address)
	log.Fatalln(http.ListenAndServe(address, nil))

}
