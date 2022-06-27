package main

import "fmt"

//https://golangbyexample.com/adapter-design-pattern-go/

//computer
type computer interface {
	insertInSquarePort()
}

//mac
type mac struct{}

func (m *mac) insertInSquarePort() {
	fmt.Println("Insert square port into mac machine")
}

//windowsAdapter
type windowsAdapter struct {
	windowMachine *windows
}

func (w *windowsAdapter) insertInSquarePort() {
	w.windowMachine.insertInCirclePort()
}

//windows
type windows struct{}

func (w *windows) insertInCirclePort() {
	fmt.Println("Insert circle port into windows machine")
}

//client
type client struct{}

func (c *client) insertSquareUsbInComputer(com computer) {
	com.insertInSquarePort()
}

func main() {
	client := &client{}
	mac := &mac{}
	client.insertSquareUsbInComputer(mac)
	windowsMachine := &windows{}
	windowsMachineAdapter := &windowsAdapter{
		windowMachine: windowsMachine,
	}
	client.insertSquareUsbInComputer(windowsMachineAdapter)
}
