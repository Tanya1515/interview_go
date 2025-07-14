package main

import (
	"errors"
	"log"
	"net"
)

// nc 127.0.0.1 12345

func main() {

	// здесь работает некоторый listener, который слушает на tcp-порту 12345
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatal(err)
	}

	for {
		// функция Accept ждет некоторого клиента, а затем возвращает connection
		// c этим клиентом. После этого вызывается функция, в рамках
		// которой отрабатывает паника. При этом упадет не одна горутина, а все приложение.
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		go ClientHandler(conn)
	}
}

func ClientHandler(c net.Conn) {
	panic(errors.New("internal error"))
}
