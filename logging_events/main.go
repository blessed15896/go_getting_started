package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	str := "abcdefghi"
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Fatalln("Cannot parse string:", err)
	}
	fmt.Println("Number is", num)
}
