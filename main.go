package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var level int
var rootDir *Folder
var cwd *Folder
var initDir string

// Folder struct type to store Folder heirarchy in tree
type Folder struct {
	Name   string
	Folder map[string]*Folder
	Path   string
}

// BackMap struct type to store immediate previous directory address
var BackMap map[*Folder]*Folder

func main() {
	log.SetFlags(log.Lshortfile)
	fmt.Println("loading application...")
	initDir = string(os.Args[1])
	initDir = checkSuffix(initDir)
	initFS()
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
