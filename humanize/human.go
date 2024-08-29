package main

import (
	"fmt"

	human "github.com/dustin/go-humanize"
)

func main() {
	var number uint64 = 123456789
	fmt.Println("Size of file is", human.Bytes(number))
}
