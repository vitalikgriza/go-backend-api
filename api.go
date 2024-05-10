package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJson(w http.ResponseWriter, status int, value any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(value)

}

type apiFunc func(http.ResponseWriter, *http.Request) error

func decorateHttpHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// handle error here
		}
	}
}

type ApiServer struct {
	listenAddr string
}

func NewServer(url string) *ApiServer {
	return &ApiServer{
		listenAddr: url,
	}
}

func (s *ApiServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", decorateHttpHandler(s.handleAccount))
	router.HandleFunc("/account/{id}", decorateHttpHandler(s.handleGetAccount))
	fmt.Println("Server running on port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func (s *ApiServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetAccount(w, r)
	case "POST":
		s.handleCreateAccount(w, r)
	case "DELETE":
		s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("method not allowd %s", r.Method)
}

func (s *ApiServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	fmt.Printf("sent id %s", id)
	account := NewAccount("Vitaly", "G")
	account.Lastname = id
	return WriteJson(w, http.StatusOK, account)
}

func (s *ApiServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
