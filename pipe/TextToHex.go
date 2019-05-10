package pipe

import (
	"fmt"
)

func TextToHex(text string) string {

	return fmt.Sprintf("%X", text) // brukte Sprint istedenfor print fordi sprint endrer outputtet til en string og ikke en int

}
