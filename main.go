package main

import (
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

func suma() float64 {
	var firstValue float64
	var secondValue float64
	fmt.Print("enter the first value: ")
	fmt.Scan(&firstValue)
	//User input of value two
	fmt.Print("enter the second value: ")
	fmt.Scan(&secondValue)
	result := firstValue + secondValue
	operator := "Addition"
	history = append(history, Operation{Op: operator, Value: result})

	return result
}

func resta() float64 {
	var firstValue float64
	var secondValue float64
	fmt.Print("enter the first value: ")
	fmt.Scan(&firstValue)
	//User input of value two
	fmt.Print("enter the second value: ")
	fmt.Scan(&secondValue)
	result := firstValue - secondValue

	operator := "Subtraction"
	history = append(history, Operation{Op: operator, Value: result})

	return result
}

func multiplicacion() float64 {
	var firstValue float64
	var secondValue float64
	fmt.Print("enter the first value: ")
	fmt.Scan(&firstValue)
	//User input of value two
	fmt.Print("enter the second value: ")
	fmt.Scan(&secondValue)
	result := firstValue * secondValue

	operator := "Multiplication"
	history = append(history, Operation{Op: operator, Value: result})

	return result
}

func division() float64 {
	var firstValue float64
	var secondValue int
	var secondValueStr string
	var secondValueFloat float64
	fmt.Print("enter the first value: ")
	fmt.Scan(&firstValue)
	//User input of value two
	for {
		//User input of value two
		fmt.Print("enter the second value: ")
		fmt.Scan(&secondValueStr)

		secondValue, _ = strconv.Atoi(secondValueStr)
		secondValueFloat = float64(secondValue)

		if secondValueFloat != 0 {
			break
		}
		fmt.Println("Second value cannot be zero. Please enter a valid number.")
	}

	result := firstValue * secondValueFloat

	operator := "Division"
	history = append(history, Operation{Op: operator, Value: result})

	return result
}

func raiz() float64 {
	var firstValue float64
	fmt.Print("enter the number: ")
	fmt.Scan(&firstValue)
	//User input of value two

	operator := "Square Root"
	history = append(history, Operation{Op: operator, Value: math.Sqrt(firstValue)})

	return math.Sqrt(firstValue)
}

func exponencial() float64 {
	var firstValue float64
	var secondValue float64
	fmt.Print("enter the numer: ")
	fmt.Scan(&firstValue)
	//User input of value two
	fmt.Print("enter the exponential: ")
	fmt.Scan(&secondValue)

	operator := "Exponential"
	history = append(history, Operation{Op: operator, Value: math.Pow(firstValue, secondValue)})

	return math.Pow(firstValue, secondValue)
}

func main() {
	//variables
	var seleccionMenu int
	today := time.Now()

	fmt.Printf("Welcome to the calculator!\nToday is day %d of month %d, %d.\n", today.Day(), int(today.Month()), today.Year())

	for {
		fmt.Println("1. Addition")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Multiplication")
		fmt.Println("4. Division")
		fmt.Println("5. Square Root")
		fmt.Println("6. Exponential")
		fmt.Println("7. History")
		fmt.Println("8. Exit Program")
		fmt.Print("Select the operation you want to perform: ")
		_, err := fmt.Scan(&seleccionMenu)
		if err != nil {
			// Si ocurre un error, imprimir mensaje de error
			fmt.Println("Invalid input. Please enter an integer.")
			// Limpiar el buffer de entrada
			// Esto asegura que cualquier dato incorrecto en el buffer sea descartado
			os.Stdin.Read(make([]byte, 1024))
			continue // Volver al inicio del loop para solicitar nuevamente la entrada
		}

		switch seleccionMenu {
		case 1:
			resultado := suma()
			fmt.Printf("Result: %.2f\n", resultado)
		case 2:
			resultado := resta()
			fmt.Printf("Resultado: %.2f\n", resultado)
		case 3:
			resultado := multiplicacion()
			fmt.Printf("Resultado: %.2f\n", resultado)
		case 4:
			resultado := division()
			fmt.Printf("Resultado: %.2f\n", resultado)
		case 5:
			resultado := raiz()
			fmt.Printf("Resultado: %.2f\n", resultado)
		case 6:
			resultado := exponencial()
			fmt.Printf("Resultado: %.2f\n", resultado)
		case 7:
			historial()
		case 8:
			fmt.Println("Saliendo del Programa...")
			os.Exit(0)
		default:
			fmt.Println("Invalid action")
		}
		fmt.Println("========================================")
		fmt.Println("================GusBunny================")
		fmt.Println("========================================")
	}

}
