# Simple Calculator in Go

This is a simple calculator program written in Go that allows you to perform basic operations like addition, subtraction, multiplication, division, square roots, exponentiation, and view the operation history.

## Features

- **Operations**: Performs the following mathematical operations:
  - Addition
  - Subtraction
  - Multiplication
  - Division
  - Square Root
  - Exponentiation
- **History**: Displays a history of all operations performed during the session.
- **Command-line Interface**: The user interacts with the program through a console menu.
- **Input Validation**: Ensures that the user enters valid values and prevents division by zero.
- **Persistent History**: Saves the operation history in memory during the program's execution.

## Requirements

- Go 1.18 or higher

## Usage Instructions

1. **Compile and Run the Program**:
    - Make sure you have Go installed on your machine. If not, download and install it from [here](https://golang.org/dl/).
    - Clone or download this repository.
    - Navigate to the folder where the Go file is located and run the following command in your terminal:
      
      ```bash
      go run main.go
      ```

2. **Interacting with the Program**:
    - Upon starting, the program will greet you with the current date and display a menu with the following options:
      - **1. Addition**
      - **2. Subtraction**
      - **3. Multiplication**
      - **4. Division**
      - **5. Square Root**
      - **6. Exponential**
      - **7. Operation History**
      - **8. Exit Program**
    - To select an operation, enter the corresponding number and follow the prompts to enter values.

3. **Operation History**:
    - The operation history is displayed by selecting option `7` from the menu. It contains all the operations performed during the session.

4. **Input Validation**:
    - The program ensures that the user's inputs are valid and prevents invalid operations, such as division by zero.

## Program Functions

### `suma()`
Adds two numbers entered by the user and displays the result.

### `resta()`
Subtracts two numbers entered by the user and displays the result.

### `multiplicacion()`
Multiplies two numbers entered by the user and displays the result.

### `division()`
Divides two numbers entered by the user, ensuring that the divisor is not zero.

### `raiz()`
Calculates the square root of a number entered by the u

