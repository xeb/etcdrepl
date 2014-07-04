package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	for true {
		fmt.Print("etcdrepl> ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		if text == "quit\n" {
			return
		}
		fmt.Println(text)
	}
}
