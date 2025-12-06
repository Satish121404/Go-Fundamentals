package main

import "fmt"

func loops() {
	var choice int
	var a, b float64

	for {
		fmt.Println("\n--- Calculator Menu ---")
		fmt.Println("1. Addition")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Multiplication")
		fmt.Println("4. Division")
		fmt.Println("5. Exit")

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		if choice == 5 {
			fmt.Println("Exiting Calculator... Goodbye!")
			break
		}

		fmt.Print("Enter first number: ")
		fmt.Scan(&a)
		fmt.Print("Enter second number: ")
		fmt.Scan(&b)

		switch choice {
		case 1:
			fmt.Printf("Result: %.2f\n", a+b)
		case 2:
			fmt.Printf("Result: %.2f\n", a-b)
		case 3:
			fmt.Printf("Result: %.2f\n", a*b)
		case 4:
			if b != 0 {
				fmt.Printf("Result: %.2f\n", a/b)
			} else {
				fmt.Println("Error: Division by zero not allowed!")
			}
		default:
			fmt.Println("Invalid choice, try again!")
		}
	}
}

func main() {
	loops()
}
