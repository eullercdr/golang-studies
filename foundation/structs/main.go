//go is not oo language
//go way
//struct is not a class

package main

import "fmt"

type Address struct {
	Street  string
	City    string
	ZipCode string
}

type Customer struct {
	Name    string
	Age     int
	Active  bool
	Address //Address is compose of Customer
}

func (customer Customer) Enable() {
	customer.Active = true
	fmt.Println("Customer is enabled")
}

func (customer Customer) Disable() {
	customer.Active = false
}

func main() {
	customer := Customer{
		Name:   "John",
		Age:    30,
		Active: true,
	}
	customer.Active = false
	customer.Address.City = "New York"
	customer.City = "New York" //other way to access
	customer.Address.Street = "Broadway"
	customer.Address.ZipCode = "10001"
	customer.Enable()
	fmt.Println(customer)
	fmt.Println(customer.Name)
	fmt.Println(customer.Age)
	fmt.Println(customer.Active)
	fmt.Println(customer.Address.City)
	fmt.Println(customer.Address.Street)
}
