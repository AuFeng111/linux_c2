package main

import (
	"bufio"
	"encoding/base32"
	"fmt"
	"os"
)

func main() {
	// Data to be encoded
	fmt.Println("example:  192.168.2.2:8888")
	fmt.Print("please enter:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	data := []byte(input)

	// Encode data to base32 string
	str := base32.StdEncoding.EncodeToString(data)

	// Print encoded string
	fmt.Println(str)
}