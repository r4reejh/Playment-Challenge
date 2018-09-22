package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var currDir string
var level int
var rootDir *Folder

// Folder struct type to store Folder heirarchy in tree
type Folder struct {
	Name   string
	Folder map[string]*Folder
}

// BackMap struct type to store immediate previous directory address
var BackMap map[*Folder]*Folder

func newFolder(name string) *Folder {
	return &Folder{name, make(map[string]*Folder)}
}

func addFolder(path string, name string, folder *Folder) bool {
	paths := deleteEmpty(strings.SplitAfter(path, "/"))
	//name := paths[len(paths)-1]
	if len(paths) > 1 {
		paths = paths[1 : len(paths)-1]
	}
	currPath := rootDir
	//fmt.Println(paths)
	for _, v := range paths {
		if currPath.Folder[v] != nil {
			currPath = currPath.Folder[v]
		} else {
			return false
		}
	}

	if currPath.Folder[name] == nil {
		currPath.Folder[name] = folder
		BackMap[folder] = currPath
	}
	return true
}

func addFolderForce(path string) bool {
	paths := deleteEmpty(strings.SplitAfter(path, "/"))
	currPath := rootDir
	for _, v := range paths {
		if currPath != nil && currPath.Folder[v] != nil {
			currPath = currPath.Folder[v]
		} else if currPath == nil {
			rootDir = newFolder(v)
			currPath = rootDir
		} else {
			folder := newFolder(v)
			currPath.Folder[v] = folder
			BackMap[folder] = currPath
			currPath = currPath.Folder[v]
		}
	}
	return true
}

func loadFileSystem(path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Print(err)
		return nil
	}
	if info.IsDir() {
		currDir = path
		folder := &Folder{info.Name() + "/", make(map[string]*Folder)}
		name := info.Name() + "/"
		addFolder(path, name, folder)
	}
	return nil
}

func indentPrint(level int) {
	for i := 0; i < level; i++ {
		fmt.Print(" ")
		/*if i == level-1 {
			fmt.Print("/")
		}*/
	}
}

func printMap(folder *Folder, spaces int) {
	indentPrint(spaces)
	fmt.Println(folder.Name)
	for _, v := range folder.Folder {
		printMap(v, spaces+1)
	}
}

func main() {
	log.SetFlags(log.Lshortfile)
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
	printMap(rootDir, 0)
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
