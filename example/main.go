package main

import (
	"fmt"
	"ofutils"
)

func main() {
	fmt.Println(ofutils.GetStructName(&CustomStruct{}))
}

type CustomStruct struct {
	Id int
}
