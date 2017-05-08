package main

import (
	"fmt"
	"net"
	"bufio"
	"brosRemote/linuxController"
)

func main(){
	//var fileName string = "things.txt"
	var port string = ":3000"
	server, _ := net.Listen("tcp", port)
	connection, _ := server.Accept()
	connection.Write([]byte("Connected\n"))
	Loop:
	for{
		input, _ := bufio.NewReader(connection).ReadString('\n')
		switch{
		case input == "quit\n":
			break Loop
		case input == "hello\n":
			connection.Write([]byte("hello " + linuxController.GetUser()))
		case input == "volumeUp\n":
			linuxController.TurnVolumeUp("10")
			fmt.Println("turning Volume up")
		case input == "volumeDown\n":
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
		}
	}
}
