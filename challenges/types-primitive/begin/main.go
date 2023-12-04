package main

import "fmt"

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val = "global"

func main() {
	// create a local variable "val" and assign it an int value
	val := 10

	// print the value of the local variable "val"
	fmt.Printf("local varibale val is %d\n", val)

	// print the value of the package-level variable "val"
	printGlobal()
	// update the package-level variable "val" and print it again
	updateGlobal()
	printGlobal()
	// print the pointer address of the local variable "val"
	fmt.Printf("memory address of val is %v\n", &val)
	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	*(&val) = 300
	fmt.Printf("Updated value of val is %v\n", val)
}

func printGlobal(){
	fmt.Printf("Value of package variable val is %s\n", val)
}

func updateGlobal(){
	val = "new Global"
}