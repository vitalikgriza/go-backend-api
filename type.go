package main

import "math/rand"

type Account struct {
	ID        int     `json:"id"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Balance   float32 `json:"balance"`
	Number    int64   `json:"accountNumber"`
}

func NewAccount(firstname, lastname string) *Account {
	return &Account{
		ID:        rand.Intn(1000),
		Firstname: firstname,
		Lastname:  lastname,
		Number:    int64(rand.Intn(100000000)),
	}
}
