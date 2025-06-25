package main

import "fmt"

type EmployeeGoogle struct {
	FirstName string
	Lastname  string
	Age       int
}
type EmployeeMeta struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// PascaleCase ( This format is used to decalre Structs , interfaces , enums)
	//  Eg. CalculateArea , UserInfo , NewHTTPRequest

	// snake_case (this format is commonly used for filename , variables)
	//  Eg. user_id , first_name , last_name

	// UPPERCASE (This format is commonly used for constants)
	// Eg. ID

	// mixedCase (This format is generally used for variables for certain external identifier)
	// Eg. javaScript , htmlDocument , isValid

	const MAXRETRIES = 5
	var employeeID = 10
	fmt.Println("employeeID " , employeeID)

}