package main

import (
	"fmt"
	"net"
	"bufio"
	
)

func main(){
	tcpServer, err := net.Listen(network: "tcp", address: ":8080")

	if err != nil{
		fmt.Println(address: "Error: ", err)
		###Idk If this "address" is correct. it was typed as "a...:"
		return
	}

	util.PrintStartUpMessage()

	defer tcpServer.Close()

	for{
		connection, _ := tcpServer.Accept()

		go handle(connection)
	}
}

func handle(connection net.Conn){

	reader := bufio.NewReader(conection)
	buffer := make([]byte, 2048)

	byteRead, _ := reader.Read(buffer)

	
}
