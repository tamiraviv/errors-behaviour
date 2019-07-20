package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"errors-behavior/errors"

	"github.com/gorilla/mux"
)

const StatusUnkownError = 520

type behaviourError interface {
	IsBadRequest() bool
	IsInternalError() bool
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	var handler http.Handler
	var handlerFunc http.HandlerFunc

	handlerFunc = myHandler
	handler = handlerFunc

	router.
		Methods("POST").
		Path("/").
		Handler(handler)

	return router
}

func main() {
	r := NewRouter()
	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	num, err := strconv.Atoi(string(body))
	if err != nil {
		log.Println("The following body is not a number:", string(body))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = isEven(num)
	if err != nil {
		if be, ok := err.(behaviourError); ok {
			if be.IsBadRequest() {
				log.Println("Got the following behaviour error - bad request:", be)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if be.IsInternalError() {
				log.Println("Got the following behaviour error - Internal server:", be)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			log.Println("Got the following behaviour error - Unkown:", be)
			w.WriteHeader(StatusUnkownError)
			return
		}

		log.Println("Got the following error:", err)
		w.WriteHeader(StatusUnkownError)
	}

	log.Printf("Yes! number %d is even!\n", num)
	w.WriteHeader(http.StatusOK)
}

func isEven(num int) (bool, error) {
	if num <= 0 {
		return false, errors.New("negative or zero number").AddBehaviour(errors.BadRequest)
	} else if num%2 == 1 {
		return false, errors.New("Invalid odd number").AddBehaviour(errors.InternalError)
	}

	return true, nil
}
