// // Package Clause - code is split into packages. Must have at least one package per GO file, but can have multiple packages per file
// // "main" was chosen as the package name as it tells the compiler that this is the main package our code is in
// package main

// // Import the fmt package, part of GO standard library, contains the Print function
// import "fmt"

// // A function in GO
// func main() {
// 	// Print function/command prints "Hello World" string to the console/terminal
// 	fmt.Print("Hello World")
// 	// In GO only double quotes can be used, single quotes cannot be used like they can in JavaScript
// }

package main

import "math"

func main() {
	// variable naming convention in GO is camelCase
	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10

	var futureValue = float64(investmentAmount) * math.Pow(1 + expectedReturnRate / 100, float64(years))
	
}
