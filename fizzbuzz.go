package main

import "fmt"

func main() {
	count := 100

	for number := 0; number <= count; number++ {

		if number%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if number%3 == 0 {
			fmt.Println("Fizz")
		} else if number%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(number)
		}
	}
}
