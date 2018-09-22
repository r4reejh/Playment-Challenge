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
	fmt.Println("loading application...")
	initFS()
	initCommands()
	fmt.Println("filesystem loaded...")
	fmt.Print("$ ")
	for true {
		var cmd string
		reader := bufio.NewReader(os.Stdin)
		cmd, _ = reader.ReadString('\n')
		processCommand(strings.TrimLeft(cmd, "\n"))
		fmt.Print("$ ")
	}
}
