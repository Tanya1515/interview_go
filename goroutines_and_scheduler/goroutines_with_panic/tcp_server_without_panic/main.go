package main

import (
	"log"
	"net"
)

// nc 127.0.0.1 12345

func main() {
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		go ClientHandler(conn)
	}
}

func ClientHandler(c net.Conn) {
	// здесь хороший пример, как обрабатывать панику и закрывается соединение при помощи defer.
	// В этом варианте паника будет обработана, и приложение не упадет. Это необходимо, поскольку
	// на одном сервере может крутится несколько клиентов, и при этом если свалится одно соединение,
	// остальные клиенты не должны получить 500 по непонятной причине.
	defer func() {
		if v := recover(); v != nil {
			log.Println("captured panic:", v)
		}
		c.Close()
	}()

	panic("internal error")
}
