package main

func main() {
	//go contains only for to loop
	for i := 0; i < 10; i++ {
		println(i)
	}
	numbers := []int{1, 2, 3, 4, 5}
	for i, n := range numbers {
		println(i, n)
	}
	for i, _ := range numbers {
		println(i)
	}
	i := 0
	for i < 10 {
		println(i)
		i++
	}
	for {
		println("Infinite loop")
	}
}
