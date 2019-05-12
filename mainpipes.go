package main

import (
	"fmt"
	"log"
	"os"

	"github.com/marcustran/is105-ica03/pipe"
)

func main() {
	data, err := os.Stat("files/text1.txt")
	if err != nil {
		log.Fatal(err)
	}
	file := data.Mode()
	fmt.Printf("info about file %v \n", data.Name)
	b := data.Size()

	fmt.Printf("Size: %v:\n", b)
	meh := string(data)
	fmt.Println("Written", len(data), "bytes \n")
	hex := pipe.TextToHex(string(data))
	fmt.Println("Written", len(hex), "bytes \n")
	fmt.Println(hex)
	jk := pipe.HexBase(hex)
	fmt.Println("Written", len(jk), "bytes \n")
	fmt.Println(jk)
	gzip := pipe.B64ToGzip(jk)
	fmt.Println("Written", len(gzip), "bytes \n")
	ALL := pipe.ALL(meh)
	fmt.Printf("Written %d bytes \n", ALL)
}
