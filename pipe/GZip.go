package pipe

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func B64ToGzip(hexToBase64 string) string {
	var buffs bytes.Buffer
	gWriter := gzip.NewWriter(&buffs)
	if _, err := gWriter.Write([]byte(hexToBase64)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	gWriter.Close()

	return fmt.Sprintln(buffs)
}
