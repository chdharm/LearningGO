package main

import (
	"fmt"
	"os"
	"bufio"
	)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(reader)
	fmt.Printf("hello, world\n")
}