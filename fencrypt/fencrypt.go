/*
Author: mzram
Date: 29/04/2023
*/

package main

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"crypto/rc4"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Configs SPECS
var (
	__CAPACITY   int = 32
	__LETTER_LEN     = 0
	__NUMBER_LEN     = 0
	__SYMBOL_LEN     = 0
)

// RULES
var (
	OPTION_RULE [3]string = [3]string{"Letter", "Number", "Symbol"}

	Letters []string = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	Numbers []string = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	Symbols []string = []string{"#", "=", "(", "}", "{", "]", ")", "[", "_", "-", "$", "|"}
)

// Rules

func __Verify_KEY_(key string) string {
	if key != "auto" {
		return key
	}

	// Create Seek(key)
	/*
		Specs:
			- 28 bytes key len
			- Random Numbers and Letters
	*/

	// Declare Params and generate seed Rule Option
	__LETTER_LEN = len(Letters)
	__NUMBER_LEN = len(Numbers)
	__SYMBOL_LEN = len(Symbols)

	data_key := make([]string, 0, __CAPACITY)

	rand.Seed(time.Now().UnixNano())
	init_rule := Generate_RULE()
	CKEY := key_gen(data_key, init_rule, __CAPACITY)

	return strings.Join(CKEY, "")
}

func Generate_RULE() string {

	random_seed := rand.Intn(3)
	//fmt.Println("The random key is: ", random_seed)

	return OPTION_RULE[random_seed]
}

func key_gen(acc_data []string, RULE string, lenght int) []string {

	if len(acc_data) == lenght {
		return acc_data
	}

	switch RULE {
	case OPTION_RULE[0]:
		// Get a rand number with __LETTER_LEN range
		random_n := rand.Intn(__LETTER_LEN - 1)
		random_letter := Letters[random_n]

		// Store the random letter
		acc_data = append(acc_data, random_letter)

		// Generate new rule
		rule := Generate_RULE()

		// Call key gen again and return the data
		key_data := key_gen(acc_data, rule, lenght)
		return key_data

	case OPTION_RULE[1]:
		// Get a rand number with __NUMBER_LEN range
		random_n := rand.Intn(__NUMBER_LEN - 1)
		random_number := Numbers[random_n]

		// Store the random number
		acc_data = append(acc_data, random_number)

		// Generate new rule
		rule := Generate_RULE()

		// Call key gen again and return
		key_data := key_gen(acc_data, rule, lenght)
		return key_data

	case OPTION_RULE[2]:
		// Get a rand number with __SYMBOL_LEN range
		random_n := rand.Intn(__SYMBOL_LEN - 1)
		random_symbol := Symbols[random_n]

		// Store the random number
		acc_data = append(acc_data, random_symbol)

		// Generate new rule
		rule := Generate_RULE()

		// Call key gen again and return
		key_data := key_gen(acc_data, rule, lenght)
		return key_data

	default:
		fmt.Println("Free!")
	}

	return nil
}

func __CHK_IFile(file string) bool {

	// First check if file exist
	if _, err := os.Stat(file); err != nil {
		return false
	}

	return true

}

/*func __EnCrypt_ADply(file string, output string) {
	//file_path := file

}*/

type IfileEncrypt struct {
	path    string
	CKey    string
	content []byte
	alg     string
}

func NewIfileEncrypt(path, Ckey string) *IfileEncrypt {

	return &IfileEncrypt{
		path:    path,
		CKey:    Ckey,
		content: make([]byte, 0),
		alg:     "",
	}
}

func (ife *IfileEncrypt) save_cont(encrypt_content []byte, Ioutput string) {
	switch Ioutput {
	case "console":
		fmt.Println(encrypt_content)

	case "file":

		var text string = ""
		for {
			reader := bufio.NewScanner(os.Stdin)
			fmt.Print("File Name ->")

			reader.Scan()

			text = reader.Text()

			//text_withext := []string{text, ".bin"}
			//text = strings.Join(text_withext, "")
			// Check file exist

			fdatacheck := __CHK_IFile(text)
			if fdatacheck == true {
				fmt.Println("[ - ] Error, this file already exist, write other file name!")
			} else {
				break
			}
		}

		// Write file in temp file in C:
		text = text + ".bin"
		f, errwritefile := os.Create("C:/temp/" + text)

		if errwritefile != nil {
			fmt.Println(errwritefile)
			panic("[ - ] Something wring when tried create the file with encrypted content!")
		}
		defer f.Close()

		_, errWrite := f.Write(encrypt_content)
		if errWrite != nil {
			panic("[ - ] Error when tried write the data!")
		}

		now := time.Now()
		day := now.Day()
		month := now.Month()
		year := now.Year()

		fmt.Printf("Log: [%d:%d:%d] | File [%s] in temp folder!\n", day, month, year, text)

		fmt.Println("[ + ] Succefully save encrypted content in: ", text)
		break

	case "hexa":
		var text string = ""
		for {
			reader := bufio.NewScanner(os.Stdin)
			fmt.Print("Hexa File Name ->")

			reader.Scan()

			text = reader.Text()

			//text_withext := []string{text, ".bin"}
			//text = strings.Join(text_withext, "")
			// Check file exist

			fdatacheck := __CHK_IFile(text)
			if fdatacheck == true {
				fmt.Println("[ - ] Error, this file already exist, write other file name!")
			} else {
				break
			}
		}

		// Write file in temp file in C:
		text = text + ".hex"
		data_hexa := make([]byte, hex.EncodedLen(len(encrypt_content)))
		hex.Encode(data_hexa, encrypt_content)
		fmt.Println("Len of encode: ", len(data_hexa))

		// decode_test
		/*decode_hexa := make([]byte, hex.DecodedLen(len(data_hexa)))
		hex.Decode(decode_hexa, data_hexa)
		fmt.Println("Normal len: ", len(decode_hexa))

		r := bytes.NewReader(decode_hexa)
		data_stream, errReader := zlib.NewReader(r)
		if errReader != nil {
			panic("[ - ] Error when decompress the core dll with zlib Alght!")
		}
		defer data_stream.Close()

		data, _ := ioutil.ReadAll(data_stream)
		fmt.Println("Original data len: ", len(data))*/
		// Other
		f, err := os.Create("C:/temp/" + text)
		if err != nil {
			fmt.Println(err)
			panic("[ - ] Something wring when tried create the file with hexa encrypted content!")
		}
		defer f.Close()

		_, errwr := f.Write(data_hexa)
		if errwr != nil {
			panic("[ - ] Error when tried write the hexadecimal data!")
		}

		now := time.Now()
		day := now.Day()
		month := now.Month()
		year := now.Year()

		fmt.Printf("Log: [%d:%d:%d] | File [%s] in C:/temp folder!\n", day, month, year, text)

		fmt.Println("[ + ] Succefully save hexadecimal encrypted content in: ", text)
		break

	}
}

func (ife *IfileEncrypt) __EnCrypt_Apply(Ioutput string, Algth string) {
	switch Algth {
	case "aes":
		block, err := aes.NewCipher([]byte(ife.CKey))
		if err != nil {
			panic("[ - ] Error when tried create NewCipher!")
		}

		gcm, errcgm := cipher.NewGCM(block)
		if errcgm != nil {
			panic("[ - ] Error when generating CGM")
		}

		nonce := make([]byte, gcm.NonceSize())
		if _, errnonce := io.ReadFull(crand.Reader, nonce); errnonce != nil {
			panic("[ - ] Nonce is invalid!")
		}

		//fmt.Println(hex.EncodeToString(gcm.Seal(nonce, nonce, ife.content, nil)))

		encrypt_content := gcm.Seal(nonce, nonce, ife.content, nil)
		// OUTPUT
		ife.save_cont(encrypt_content, Ioutput)
		break

	case "rc4":
		cipher, err := rc4.NewCipher([]byte(ife.CKey))
		if err != nil {
			panic("[ - ] Error in NewCypher rc4")
		}

		cipher_text := make([]byte, len(ife.content))

		cipher.XORKeyStream(cipher_text, ife.content)

		ife.save_cont(cipher_text, Ioutput)
		break
	case "none":
		// None encriptation, Only transform to hexa and save

		// Only compress mode
		var b bytes.Buffer
		data_writer := zlib.NewWriter(&b)
		data_writer.Write(ife.content)
		data_writer.Close()

		ife.save_cont(b.Bytes(), Ioutput)

		break
	}

}

func (ife *IfileEncrypt) Instance() {
	cont, err := ioutil.ReadFile(ife.path)
	if err != nil {
		panic("[ - ] Error when tried read the File! Please check if file is correct")
	}

	ife.content = append(ife.content, cont...)

	fmt.Println("[ + ] The file have been read!")
	//println(hex.EncodeToString(ife.content))
}

func main() {

	fmt.Println("[ + ] Running IfileEncryptor By mzram")

	// Get Arguments: FileToEncypt, Input(File, Output), key

	if len(os.Args) < 4 {
		panic("[ - ] Error! Missing IFile, IOutput or IKey parameters!")
	}

	// Get parameters
	IFile := os.Args[1]
	IOutput := os.Args[2]
	IKey := os.Args[3]

	CKEY := __Verify_KEY_(IKey)
	fmt.Println("[ + ] Stablish current KEY: ", CKEY)

	// Read File and encrypt with CKEY and RC4 algothrim
	filechk := __CHK_IFile(IFile)
	if filechk == false {
		panic("[ - ] Error! This file dont exist in the path")
	}

	if IOutput != "console" && IOutput != "file" && IOutput != "hexa" {
		panic("[ - ] This Output type is not especified!")
	}

	algth := os.Args[4]
	__ENCR_Object := NewIfileEncrypt(IFile, CKEY)
	__ENCR_Object.Instance()
	__ENCR_Object.__EnCrypt_Apply(IOutput, algth) // CORE of the program

}
