package main

import (
	"bufio"
	"log"
	"os"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func GBKFileToUtf8(filePath string) string {
	// Read UTF-8 from a GBK encoded file.
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	r := transform.NewReader(f, simplifiedchinese.GBK.NewDecoder())

	// Read converted UTF-8 from `r` as needed.
	// As an example we'll read line-by-line showing what was read:
	sc := bufio.NewScanner(r)
	result := ""
	for sc.Scan() {
		result += string(sc.Bytes()) + "\n"
	}
	return result
}
