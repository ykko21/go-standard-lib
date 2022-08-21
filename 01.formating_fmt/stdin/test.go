package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("First name:")
	s, _ := reader.ReadString('\n')
	fmt.Println(s)
}
