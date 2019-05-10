package pipe

import (
	"compress/gzip"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
)

func ALL(translate string) string {
	r := gzip.NewWriter(base64.NewEncoder(base64.StdEncoding, hex.NewEncoder(os.Stdout)))
	r.Write([]byte(translate))
	r.Close()

	return fmt.Sprintln(translate)
}
