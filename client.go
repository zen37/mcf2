package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const url string = "http://localhost:8080/v1/payments"

//curl localhost:1111/v1/payments  -d '{"ID":"123"}'
func main() {

	fmt.Println("URL:>", url)

	requestBody, err := json.Marshal(map[string]string{"ID": "hjdha2123h3"})

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))

	/*
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))


			req.Header.Set("X-Custom-Header", "myvalue")
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("User-Agent", "my user agent")

			client := &http.Client{}
			resp, err := client.Do(req)
	*/
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
