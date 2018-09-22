package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var rootDir *Folder
var cwd *Folder
var initDir string

func main() {
	fmt.Println("$> <Starting your application...>")
	initFS()
	initCommands()
	fmt.Print("$> ")
	for true {
		var cmd string
		reader := bufio.NewReader(os.Stdin)
		cmd, _ = reader.ReadString('\n')
		if string(strings.TrimRight(cmd, "\n")) == "exit" {
			fmt.Println("Bye!")
			return
		}
		processCommand(strings.TrimRight(cmd, "\n"))
		fmt.Print("$> ")
	}
}
