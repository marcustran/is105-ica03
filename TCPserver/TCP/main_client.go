package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type User struct {
	Name string
	Mail string
}

func main() {

	conn, err := net.Dial("tcp", "10.228.17.13:9000")
	if err != nil {
		log.Fatal(err)
	}

	b, err := bufio.NewReader(conn).ReadBytes(byte('\n'))
	var u User
	err = json.Unmarshal(b, &u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(u.Name)
	fmt.Println(u.Mail)
}
