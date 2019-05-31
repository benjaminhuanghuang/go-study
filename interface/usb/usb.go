package main

import "fmt"

type Connecter interface {
	Connect()
}

type USB interface {
	Name() string
	Connecter
}

type PhoneConnecter struct {
	name string
}

// Implement interface
func (phoneConn PhoneConnecter) Name() string {
	return phoneConn.name
}

func (phoneConn PhoneConnecter) Connect() {
	fmt.Println("connect:", phoneConn.name)
}

func main() {
	var usb USB
	usb = PhoneConnecter{"MyUSB"}
	usb.Connect()
}
