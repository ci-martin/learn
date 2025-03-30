package app

import "fmt"

func init() {
	// Initialization code here, if needed
	fmt.Println("Initializing the app package...")
	app()
}

func app() {

	// Take user input
	fmt.Println("Please enter your name:")
	var name string
	fmt.Scanln(&name)

	// Print a greeting message
	fmt.Printf("Hello, %s! Welcome to the application.\n", name)
}
