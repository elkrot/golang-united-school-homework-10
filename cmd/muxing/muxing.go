package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

GET		/name/{PARAM}						body: Hello, PARAM!
GET		/bad								Status: 500
POST	/data + Body PARAM					body: I got message:\nPARAM
POST	/headers+ Headers{"a":"2", "b":"3"}	Header "a+b": "5"

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name/{param}",ParameterListHandler).Methods(http.MethodGet)
	router.HandleFunc("/bad",BadResponseHandler).Methods(http.MethodGet)
	router.HandleFunc("/data",DataHandler).Methods(http.MethodPost)
	router.HandleFunc("/headers",HeaderHandler).Methods(http.MethodPost)
/*	router.HandleFunc("/api/v0/user",user.List).Methods(http.MethodGet)
	router.HandleFunc("/api/v0/user",user.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/v0/user/{id}",user.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/api/v0/user/{id}",user.View).Methods(http.MethodGet)
	router.HandleFunc("/api/v0/user/{id}",user.Update).Methods(http.MethodPut)*/

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))

	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}
func BadResponseHandler(w http.ResponseWriter, r *http.Request) {    
    w.WriteHeader(http.StatusOK)   
}
func DataHandler(w http.ResponseWriter, r *http.Request) {  
	vars := mux.Vars(r)  
    w.WriteHeader(http.StatusOK)   
	fmt.Fprintf(w, "I got message:\n%v!\n", vars["param"])
}
func HeaderHandler(w http.ResponseWriter, r *http.Request) { 
	vars := mux.Vars(r)   
	r.Headers("X-Requested-With", "XMLHttpRequest")
    w.WriteHeader(http.StatusOK)   
}
func ParameterListHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Hello, %v!\n", vars["param"])
}



//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
