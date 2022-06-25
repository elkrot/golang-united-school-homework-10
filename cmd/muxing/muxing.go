package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name/{param}", ParameterListHandler).Methods(http.MethodGet)
	router.HandleFunc("/bad", BadResponseHandler).Methods(http.MethodGet)
	router.HandleFunc("/data", DataHandler).Methods(http.MethodPost)
	router.HandleFunc("/headers", HeaderHandler).Methods(http.MethodPost)
	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}
func BadResponseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}
func DataHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "I got message:\n%v", string(b))
}
func HeaderHandler(w http.ResponseWriter, r *http.Request) {

	a, _ := strconv.Atoi(r.Header.Get("a"))
	b, _ := strconv.Atoi(r.Header.Get("b"))
	a_b := a + b
	w.Header().Set("a+b", fmt.Sprint(a_b))
	w.WriteHeader(http.StatusOK)
}
func ParameterListHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %v!", vars["param"])
}

func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
