package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/marcustran/is105-ica03/pipe"
)

func main() {
	data, err := ioutil.ReadFile("files/pg100.txt")
	if err != nil {
		log.Fatal(err)
	}
	hex := pipe.TextToHex(string(data))
	gzip := pipe.B64ToGzip(hex)
	fmt.Printf("Written %d bytes \n", gzip)
	ALL := pipe.ALL
	fmt.Printf("Written %d bytes \n", ALL)
}
