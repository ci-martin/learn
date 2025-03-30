// File: message.go
// Project: Go Multi-file Project
// Description: This file contains the Message struct and methods to create and send messages.

package main

import "time"

type Message struct {
	MessageID uint32
	Message   string
	TimeStamp time.Time
	Recipient string
	Sender    string
}

type Sender struct {
	Name string
}

type Recipient struct {
	Name string
}

// NewMessage creates a new message with the given details.
func NewMessage(message string, recipientName string, senderName string) *Message {

	// Create a new recipient and sender
	thisRecipient := Recipient{}
	thisRecipient.NewRecipient(recipientName)

	thisSender := Sender{}
	thisSender.NewSender(senderName)

	// Create a new message with the given details
	msg := &Message{
		MessageID: generateMessageID(),
		Message:   message,
		TimeStamp: time.Now(),
		Recipient: thisRecipient.Name,
		Sender:    thisSender.Name,
	}

	return msg
}

func generateMessageID() uint32 {
	// Generate a unique message ID (for simplicity, using a static value here)
	return 123456
}

func (Recipient *Recipient) NewRecipient(name string) {
	// Create a new recipient with the given name
	Recipient.Name = name
}

func (Sender *Sender) NewSender(name string) {
	// Create a new sender with the given name
	Sender.Name = name
}
