package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Account struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Saldo int    `json:"saldo"`
}

func main() {
	account := Account{
		ID:    "1",
		Name:  "John Doe",
		Email: "niev@ga.lk",
	}
	res, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		panic(err)
	}

	//transform json to struct
	jsonPayload := `{"ID":"2","Name":"Jane Doe","Email":"rismed@sab.bi"}`
	var account2 Account
	err = json.Unmarshal([]byte(jsonPayload), &account2)
	if err != nil {
		panic(err)
	}
	fmt.Println(account2.Email)
}
