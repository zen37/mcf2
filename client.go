package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const url string = "http://localhost:8080/v2/payments"

func main() {

	fmt.Println("http.Get: ", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("response: ", string(body))

}
