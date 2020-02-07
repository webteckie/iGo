package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	DisplayNetCoreEnvironment()
	CleanTheDatabase()
	StartTheTimer()
	StartTheReportingServices()
	WaitForExit()
}

// DisplayNetCoreEnvironment Scan user input
func DisplayNetCoreEnvironment() {
	fmt.Printf("***Running with ASPNETCORE_ENVIRONMENT = %s***\n", os.Getenv("ASPNETCORE_ENVIRONMENT"))
}

// CleanTheDatabase asks the user if the db should be cleaned and cleans it, if so
func CleanTheDatabase() {
	var startTimer string = ScanUserKeyInput("\nDo you want the database to be cleaned:", []byte(`yYnN`), "n")
	if startTimer == "y" || startTimer == "Y" {
		fmt.Println("Cleaning the database...")
	}
}

// StartTheTimer asks the user if the timer should be started and starts it, if so
func StartTheTimer() {
	var startTimer string = ScanUserKeyInput("\nDo you want the timer service to be started:", []byte(`yYnN`), "n")
	if startTimer == "y" || startTimer == "Y" {
		fmt.Println("Starting the timer...")
	}
}

// StartTheReportingServices starts the reporting services
func StartTheReportingServices() {
	fmt.Println("\nStarting the reporting services...\n")
}

// WaitForExit waits for the console test app to exit
func WaitForExit() {
	ScanUserKeyInput("\nPress x or X to exit", []byte(`xX`), "")
}

// ScanUserKeyInput scans user input
func ScanUserKeyInput(prompt string, acceptedInputs []byte, defaultAnswer string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(prompt)
	var answer string = defaultAnswer
	for scanner.Scan() {
		input := scanner.Bytes()
		if len(input) != 0 && bytes.Contains(acceptedInputs, input) != false {
			answer = string(input)
			break
		}
	}
	return answer
}
