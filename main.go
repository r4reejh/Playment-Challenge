package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var currDir string
var level int
var rootDir *Folder
var cwd *Folder

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
	dir := string(os.Args[1])
	if !strings.HasSuffix(dir, "/") {
		dir = dir + "/"
	}
	BackMap = make(map[*Folder]*Folder)
	addFolderForce(dir)
	//rootDir = newFolder("/")
	err := filepath.Walk(dir, loadFileSystem)
	if err != nil {
		log.Fatal(err)
	}
	cwd = rootDir
	fmt.Println("filesystem loaded...")
	fmt.Print("$ ")
	for true {
		var cmd string
		reader := bufio.NewReader(os.Stdin)
		cmd, _ = reader.ReadString('\n')
		processCommand(strings.TrimLeft(cmd, "\n"))
		fmt.Print("$ ")
	}
	//printMap(rootDir, 0)
}
