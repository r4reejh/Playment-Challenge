package main

import (
	"fmt"
	"strings"
)

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

func printMap(folder *Folder, spaces int) {
	indentPrint(spaces)
	fmt.Println(folder.Name)
	for _, v := range folder.Folder {
		printMap(v, spaces+1)
	}
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
