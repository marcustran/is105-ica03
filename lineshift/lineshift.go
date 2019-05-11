package main

import ()
	"github.com/arnekd/ICA02/fileutils/"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func DetermineLineshift(bSlice []byte) bool {

	for _, b := range bSlice {
		if b == 13 {
			return true
		}
	}
	return false
}

func main() {

	if len(os.Args) == 1 {
		return
	}

	bSlice := fileutils.FileToByteslice(os.Args[1])

	windows := DetermineLineshift(bSlice)

	if windows {
		fmt.Println("Windows file")
	} else {
		fmt.Println("Unix/Mac file")
	}

}
