package main

import "fmt"

func main() {
	fmt.Println("Maps")
	salaries := map[string]int{
		"john": 1000,
		"jane": 1500,
		"joe":  800,
	}
	delete(salaries, "jane")
	salaries["jane"] = 2000
	fmt.Println(salaries["jane"])

	sal := make(map[string]int)
	sal["johny"] = 1000

	for name, salary := range salaries {
		fmt.Printf("%s has salary %d\n", name, salary)
	}

	for _, salary := range salaries {
		fmt.Printf("Salary: %d\n", salary)
	}
}
