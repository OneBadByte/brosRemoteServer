package main

import (
	"brosRemote/linuxController"
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
	//var fileName string = "things.txt"
	var port string = ":3000"
	server, _ := net.Listen("tcp", port)
	connection, _ := server.Accept()
	connection.Write([]byte("Connected\n"))
Loop:
	for {
		input, _ := bufio.NewReader(connection).ReadString('\n')
		switch {
		case input == "quit\n":
			break Loop
		case input == "hello\n":
			connection.Write([]byte("hello " + linuxController.GetUser()))
		case input == "up\n":
			linuxController.TurnVolumeUp("10")
			fmt.Println("turning Volume up")
		case input == "down\n":
			linuxController.TurnVolumeDown("10")
			fmt.Println("turning Volume down")
		case input == "mute\n":
			linuxController.MuteVolume(true)
			fmt.Println("muting")
		case input == "unmute\n":
			linuxController.MuteVolume(false)
			fmt.Println("unmuting")
		case input == "lock\n":
			linuxController.TriggerLockScreen()
		case input == "restart\n":
			linuxController.RestartComputer("^6^Linux^6^\n")
		case input == "update\n":
			linuxController.UpdateComputer("^6^Linux^6^\n")
		case input == "shutdown\n":
			linuxController.ShutDownComputer("^6^Linux^6^\n")
		case input == "command\n":
			input, _ := bufio.NewReader(connection).ReadString('\n')
			stuff := RemoveNewLine(input)
			output, err := linuxController.RunCommand(stuff)
			fmt.Println(output, err)


		}
	}
}
