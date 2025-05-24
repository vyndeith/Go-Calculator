# Go Calculator

A simple, robust command-line calculator built in Go that performs basic arithmetic operations with input validation and error handling.

## Features

- **Basic Arithmetic Operations**: Addition (+), Subtraction (-), Multiplication (*), Division (/)
- **Input Validation**: Robust error handling for non-numeric inputs
- **Division by Zero Protection**: Prevents crashes from division by zero operations
- **Floating Point Support**: Works with decimal numbers for precise calculations
- **User-Friendly Interface**: Clear prompts and error messages

## Technologies Used

- **Go** (Golang)
- **Standard Library Packages**:
  - `fmt` - for formatted I/O operations
  - `strconv` - for string to number conversion

## Prerequisites

- Go 1.16 or higher installed on your system
- Basic understanding of command-line interface

## Installation & Usage

1. **Clone the repository**:
   ```bash
   git clone https://github.com/vyndeith/Go-Calculator.git
   cd Go-Calculator
   ```

2. **Run the calculator**:
   ```bash
   go run main.go
   ```

3. **Follow the prompts**:
   - Enter first number
   - Enter second number  
   - Select operation (+, -, *, /)
   - View the result

## Example Usage

```
First number: 15.5
Second number: 3.2
Select operation (+, -, *, /): *
You wrote: 15.500000 * 3.200000
Result: 49.600000
```

## Error Handling

The calculator handles various error scenarios:

- **Invalid Number Input**: Prompts user to re-enter if non-numeric input is provided
- **Division by Zero**: Displays error message and safely exits
- **Invalid Operation**: Shows "Unknown operation" message for unsupported operations

## Code Structure

- **Input Validation Function**: `inputNumber()` - Reusable function for safe number input
- **Main Logic**: Handles operation selection and calculation
- **Error Handling**: Comprehensive validation throughout the application

## Learning Outcomes

This project demonstrates:
- Go syntax and basic programming concepts
- Function creation and usage
- Error handling and input validation
- Working with different data types (`float64`, `string`)
- Control structures (`if/else`, `for` loops)
- Standard library usage
