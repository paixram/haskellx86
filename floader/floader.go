package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/*
Author: mzram
Date: 29/04/2023
*/

func main() {

	if len(os.Args) < 1 {
		panic("[ - ] Please, especify path of the file .hex!")
	}
	filepath := os.Args[1]
	myfile, _ := os.Stat(filepath)

	fmt.Println("The size of file is: ", myfile.Size())

	sizeperpacket := 1960
	sizefile := myfile.Size() / int64(sizeperpacket)

	fmt.Println(sizefile)

	// leer el archivo
	content, errf := ioutil.ReadFile(filepath)
	if errf != nil {
		fmt.Println("Error while reading file")
	}

	dest_malware_path := "C:\\temp\\mpfldr.cpp"

	instance_config := fmt.Sprintf("#include <string>\n#include <windows.h>\nusing namespace std;\n")

	code := ""
	checksum_custom := "string pld = "
	//init_count := 0
	var init_count = 0
	for i := int(sizeperpacket); ; i += int(sizeperpacket) {

		// Con buffer de 460 vamos a analizar si la posicion actual es igual a la longitud de content, si es asi, solo rellenar con 0x00

		if init_count >= len(content) && i >= len(content) {

			last_pos := strings.LastIndex(checksum_custom, "+")

			checksum_custom = checksum_custom[:last_pos] + strings.Replace(checksum_custom[last_pos:], "+", ";", 1)

			break
		}

		if init_count < len(content) && i > len(content) {
			for {
				if i == len(content) {
					break
				}

				i = i - 1
				continue
			}
		}

		first := content[init_count:i]
		unicodeID := fmt.Sprintf("unicode_%d", init_count)
		thestring := fmt.Sprintf("string %s = \"%s\";\n", unicodeID, first)

		code = code + thestring

		checksum_custom = checksum_custom + unicodeID + "+"

		fmt.Println(thestring)

		init_count += sizeperpacket
	}

	code = instance_config + code + checksum_custom
	err := ioutil.WriteFile(dest_malware_path, []byte(code), 0777)
	if err != nil {
		fmt.Println("Error al escribir en el archivo")
	}
}
