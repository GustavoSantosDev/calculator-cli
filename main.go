package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

type Operation struct {
	Op    string
	Value float64
}

var history []Operation

func historial() {
	if len(history) == 0 {
		fmt.Println("No history available.")
		return
	}

	fmt.Println("History of operations:")
	for i, op := range history {
		fmt.Printf("%d. %s: %.2f\n", i+1, op.Op, op.Value)
	}
}

func getFloatInput(prompt string) float64 {
	scanner := bufio.NewScanner(os.Stdin)
	var value float64

	for {
		fmt.Print(prompt)
		if scanner.Scan() {
			input := scanner.Text()
			var err error
			value, err = strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
			} else {
				break
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
			break
		}
	}
	return value
}

func suma() float64 {
	firstValue := getFloatInput("Enter the first value: ")
	//User input of value two
	secondValue := getFloatInput("Enter the second value: ")

	result := firstValue + secondValue
	operator := "Addition"
	history = append(history, Operation{Op: operator, Value: result})

	fmt.Printf("Result: %.2f\n", result)
	return result
}

func resta() float64 {
	firstValue := getFloatInput("Enter the first value: ")

	//User input of value two
	secondValue := getFloatInput("enter the second value: ")

	result := firstValue - secondValue
	operator := "Subtraction"
	history = append(history, Operation{Op: operator, Value: result})

	fmt.Printf("Result: %.2f\n", result)
	return result
}

func multiplicacion() float64 {
	firstValue := getFloatInput("Enter the first value: ")
	//User input of value two
	secondValue := getFloatInput("Enter the second value")

	result := firstValue * secondValue
	operator := "Multiplication"
	history = append(history, Operation{Op: operator, Value: result})

	fmt.Printf("Result: %.2f\n", result)
	return result
}

func division() float64 {
	firstValue := getFloatInput("Enter the first value: ")
	var secondValue float64
	//User input of value two
	for {
		secondValue = getFloatInput("Enter the second value: ")

		if secondValue != 0 {
			break
		}
		fmt.Println("The second value cannot be zero because division by zero is not possible. Please enter a different number to proceed.")
	}

	result := firstValue / secondValue

	operator := "Division"
	history = append(history, Operation{Op: operator, Value: result})

	fmt.Printf("Result: %.2f\n", result)
	return result
}

func raiz() float64 {
	firstValue := getFloatInput("enter the number: ")

	operator := "Square Root"
	history = append(history, Operation{Op: operator, Value: math.Sqrt(firstValue)})

	fmt.Printf("Result: %.2f\n", math.Sqrt(firstValue))
	return math.Sqrt(firstValue)
}

func exponencial() float64 {
	firstValue := getFloatInput("enter the numer: ")
	//User input of value two
	secondValue := getFloatInput("enter the exponential: ")

	operator := "Exponential"
	history = append(history, Operation{Op: operator, Value: math.Pow(firstValue, secondValue)})

	fmt.Printf("Result: %.2f\n", math.Pow(firstValue, secondValue))
	return math.Pow(firstValue, secondValue)
}

func menu() {
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Println("5. Square Root")
	fmt.Println("6. Exponential")
	fmt.Println("7. History")
	fmt.Println("8. Exit Program")

	seleccionMenu := getFloatInput("Select the operation you want to perform: ")
	fmt.Printf("You Entered: %2.f\n", seleccionMenu)

	switch seleccionMenu {
	case 1:
		suma()
	case 2:
		resta()
	case 3:
		multiplicacion()
	case 4:
		division()
	case 5:
		raiz()
	case 6:
		exponencial()
	case 7:
		historial()
	case 8:
		fmt.Println("Exiting the program...")
		os.Exit(0)
	default:
		fmt.Println("Invalid action")
	}
	fmt.Println("========================================")
	fmt.Println("================GusBunny================")
	fmt.Println("========================================")
}

func main() {
	today := time.Now()
	fmt.Printf("Welcome to the calculator!\nToday is day %d of month %d, %d.\n", today.Day(), int(today.Month()), today.Year())

	for {
		menu()
	}

}
