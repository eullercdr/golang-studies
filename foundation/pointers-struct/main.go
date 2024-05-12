package main

import "fmt"

type Account struct {
	name   string
	amount int
}

// *Account is a pointer to Account
func (c *Account) addValue(amount int) int {
	c.amount += amount
	return c.amount
}

func NewAccount(name string) *Account {
	return &Account{
		name:   name,
		amount: 100,
	}
}

func main() {
	account := NewAccount("John")
	fmt.Println(account.addValue(100))
	fmt.Println(account)

	account2 := NewAccount("Jane")
	fmt.Println(account2.addValue(200))
	fmt.Println(account2)
}
