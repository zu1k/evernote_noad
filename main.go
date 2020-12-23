package main

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os"
)

var adsString = "/third/marketing/getMarketingPromotion"

func main() {
	fileName := "Evernote.exe"
	if len(os.Args)==2 {
		fileName = os.Args[1]
	}

	utf16Encoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder()
	sig,_,_ := transform.Bytes(utf16Encoder, []byte(adsString))

	file,err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	sigIdx := bytes.Index(file, sig)
	if sigIdx>0 {
		fmt.Println("find sig:", sigIdx)
	}

	newBytes := make([]byte, len(sig))
	newBytes[0] = byte('/')
	file = bytes.ReplaceAll(file, sig, newBytes)

	err = ioutil.WriteFile(fileName, file, os.ModePerm)
	if err != nil {
		panic(err)
	}
}