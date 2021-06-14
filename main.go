package main

import (
	"fmt"
	"log"
	"net/http"
)

//health check:  https://.../mcf2/healthcheck

//payments end point : https://.../mcf2/v1.0/payments

const address string = ":8080"

func payments(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(w, r.URL.Path)
	fmt.Fprintf(w, "%s s-print %q q-print %v v-print \n", "Hello World", "Hallo Welt", "Hola Mundo")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "r.RequestURI is: %s\n", r.RequestURI)
	fmt.Fprintf(w, "r.URL.Path is: %s\n", r.URL.Path)
	fmt.Fprintf(w, "r.URL is: %s\n", r.URL)
	fmt.Fprintf(w, "r.URL.User is: %s\n", r.URL.User)
	fmt.Fprintf(w, "r.URL.Host is: %s\n", r.URL.Host)
	fmt.Fprintf(w, "r.URL.Fragment is: %s\n", r.URL.Fragment)
}

func payments2(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(w, r.URL.Path)
	fmt.Fprintf(w, "200")
}

func main() {

	http.HandleFunc("/v1/payments/", payments)

	http.HandleFunc("/v2/payments/", payments2)

	log.Fatalln(http.ListenAndServe(address, nil))

}
