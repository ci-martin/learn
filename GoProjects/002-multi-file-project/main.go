// File: main.go
// Project: Go Multi-file Project
// Description: This is the main file of a multi-file Go project that sends messages.

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, please create your message!")
	// Create a new message
	message := NewMessage("Hello, how are you?", "Alice", "Bob")

	// Print the message details
	fmt.Printf("Message sent to %s\n", message.Recipient)
	fmt.Printf("Message from %s\n", message.Sender)
	fmt.Printf("Message ID: %d\n", message.MessageID)
	fmt.Printf("Message: %s\n", message.Message)
	fmt.Printf("Timestamp: %s\n", message.TimeStamp.Format(time.RFC1123))
	fmt.Println("Message sent successfully!")
}
