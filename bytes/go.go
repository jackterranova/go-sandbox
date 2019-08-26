package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"

	"encoding/base64"
	"fmt"
	"strconv"
	//"os"
)

//make this a constant!!!
var ALPHABET = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-', '_'}

func main() {

	//key := "loyWi-DrNvJBguwp1-XXfZt1frFvT2SyDXl2WvFwa-k"
	key := "FmiMZoMWRypTPeSU1Btlr8oBF4rjpSwjJaM3UadUMiI"
	decodedkey, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(key)
	fmt.Println(err)
	//decodedkey = []byte{-106, -116, -106, -117, -32, -21, 54, -14, 65, -126, -20, 41, -41, -27, -41, 125, -101, 117, 126, -79, 111, 79, 100, -78, 13, 121, 118, 90, -15, 112, 107, -23}

	fmt.Println("DECODEDKEY ", string(decodedkey))

	data := "admin"
	ts := strconv.Itoa(1562769634781)
	unenctoken := data + "|" + ts + "|"
	fmt.Println("unenctoken: " + unenctoken)

	//hmac
	h := hmac.New(sha1.New, decodedkey)
	h.Write([]byte(unenctoken))
	standard_hmac := h.Sum(nil)
	fmt.Println("standard_hmac: ", standard_hmac)

	//keyhash
	s1 := sha1.New()
	s1.Write(decodedkey)
	key_hash := s1.Sum(nil)
	fmt.Println("key ", decodedkey)
	fmt.Println("key_hash ", key_hash)

	var b bytes.Buffer

	// add the unencode token
	b.WriteString(unenctoken)

	// calc signiture
	var signiturebuf bytes.Buffer
	signiturebuf.WriteByte(0)
	signiturebuf.Write(key_hash[0:3])
	signiturebuf.Write(standard_hmac)

	encodedsig := encodeWebSafe(signiturebuf.Bytes())

	b.WriteString(encodedsig)

	fmt.Println(encodeWebSafe(b.Bytes()))
}

func encodeWebSafe(input []byte) string {
	fmt.Println("EncodeWebSafe IN => ", input)
	inputBlocks := len(input) / 3
	remainder := len(input) % 3
	outputLen := inputBlocks * 4

	switch remainder {
	case 1:
		outputLen += 2
		break
	case 2:
		outputLen += 3
		break
	}

	var out = make([]byte, outputLen)
	outPos := 0
	inPos := 0

	for i := 0; i < inputBlocks; i++ {

		buffer := (255 & input[inPos]) << 16
		inPos++
		buffer |= (0xFF & input[inPos]) << 8
		inPos++
		buffer |= 0xFF & input[inPos]
		inPos++

		out[outPos] = ALPHABET[(buffer>>18)&0x3F]
		outPos++
		out[outPos] = ALPHABET[(buffer>>12)&0x3F]
		outPos++
		out[outPos] = ALPHABET[(buffer>>6)&0x3F]
		outPos++
		out[outPos] = ALPHABET[buffer&0x3F]
		outPos++
	}

	if remainder > 0 {
		buffer := (0xFF & input[inPos]) << 16
		inPos++

		if remainder == 2 {
			buffer |= (0xFF & input[inPos]) << 8
			inPos++
		}
		out[outPos] = ALPHABET[(buffer>>18)&0x3F]
		outPos++
		out[outPos] = ALPHABET[(buffer>>12)&0x3F]
		outPos++
		if remainder == 2 {
			out[outPos] = ALPHABET[(buffer>>6)&0x3F]
			outPos++
		}
	}

	fmt.Println("EncodeWebSafe OUT => ", out)
	return string(out)
}
