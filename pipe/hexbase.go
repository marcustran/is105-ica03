package pipe

import "encoding/base64"

func HexBase(data string) string {
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))

	return sEnc
}
