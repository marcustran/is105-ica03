package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/marcustran/is105-ica03/pipe"
)

func main() {
	data, err := ioutil.ReadFile("files/text4.txt")
	if err != nil {
		log.Fatal(err)
	}
	hex := pipe.TextToHex(string(data))
	fmt.Println("Written", len(hex), "bytes \n")
	jk := pipe.HexBase(hex)
	fmt.Println("Written", len(jk), "bytes \n")
	gzip := pipe.B64ToGzip(jk)
	fmt.Println("Written", len(gzip), "bytes \n")
	//ALL := pipe.ALL(data)
}
