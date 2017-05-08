package main

import (
	"brosRemote/linuxController"
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
	server, _ := net.Listen("tcp", port)
	connection, _ := server.Accept()
	connection.Write([]byte("Connected\n"))
Loop:
	for {
		input, _ := bufio.NewReader(connection).ReadString('\n')
		switch {
		case input == "quit\n":
			connection.Write([]byte("quitting\n"))
			break Loop
		case input == "hello\n":
			connection.Write([]byte("hello " + linuxController.GetUser()))
		case input == "up\n":
			linuxController.TurnVolumeUp("10")
			connection.Write([]byte("Turning volume up\n"))
		case input == "down\n":
			linuxController.TurnVolumeDown("10")
			connection.Write([]byte("Turning volume down\n"))
		case input == "mute\n":
			linuxController.MuteVolume(true)
			connection.Write([]byte("Muting\n"))
		case input == "unmute\n":
			linuxController.MuteVolume(false)
			connection.Write([]byte("Unmuting\n"))
		case input == "lock\n":
			linuxController.TriggerLockScreen()
			connection.Write([]byte("Locking Screen\n"))
		case input == "restart\n":
			linuxController.RestartComputer(fileController.ReadFile(passwordFile))
			connection.Write([]byte("Restarting\n"))
		case input == "update\n":
			linuxController.UpdateComputer(fileController.ReadFile(passwordFile))
			connection.Write([]byte("Updating computer\n"))
		case input == "shutdown\n":
			linuxController.ShutDownComputer(fileController.ReadFile(passwordFile))
			connection.Write([]byte("Shutting down computer\n"))
		case input == "command\n":
			input, _ := bufio.NewReader(connection).ReadString('\n')
			stuff := RemoveNewLine(input)
			output, err := linuxController.RunCommand(stuff)
			fmt.Println(output, err)


		}
	}
}
