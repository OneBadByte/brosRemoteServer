package main

import (
	"brosRemote/osController"
	"brosRemote/fileController"
	"bufio"
	"fmt"
	"net"
)

func RemoveNewLine(input string)string{
	s := len(input)
	input = string(input[:s-1])
	return input
}

func main() {
	var passwordFile string = "NothingToSeeHere.txt"
	fileController.CreateFile(passwordFile)
	var port string = ":3000"
	fmt.Println("started")
	server, _ := net.Listen("tcp", port)
	connection, _ := server.Accept()
	connection.Write([]byte("Connected\n"))
Loop:
	for {
		input, _ := bufio.NewReader(connection).ReadString('\n')
		command := RemoveNewLine(input)
		message := osController.RunCommandFromDatabase(command)
		connection.Write([]byte(message + "\n"))
		switch {
		case input == "quit\n" || input == "exit\n":
			connection.Write([]byte("quitting\n"))
			break Loop
		case input == "restart\n":
			osController.RestartComputer(fileController.ReadFile(passwordFile))
			connection.Write([]byte("Restarting\n"))
		case input == "update\n":
			osController.UpdateComputer(fileController.ReadFile(passwordFile))
			connection.Write([]byte("Updating computer\n"))
		case input == "shutdown\n":
			osController.ShutDownComputer(fileController.ReadFile(passwordFile))
			connection.Write([]byte("Shutting down computer\n"))
		case input == "command\n":
			input, _ := bufio.NewReader(connection).ReadString('\n')
			stuff := RemoveNewLine(input)
			output, err := osController.RunCommand(stuff)
			fmt.Println(output, err)


		}
	}
}
