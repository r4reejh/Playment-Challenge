package main

import (
	"fmt"
	"strings"
)

func newFolder(name string, path string) *Folder {
	return &Folder{name, make(map[string]*Folder), path}
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

	if currPath.Folder[checkSuffix(name)] == nil {
		currPath.Folder[checkSuffix(name)] = folder
		BackMap[folder] = currPath
	}
	return true
}

func addFolderForce(path string) bool {
	paths := deleteEmpty(strings.SplitAfter(path, "/"))
	currPath := rootDir
	cumulativePaths := ""
	for _, v := range paths {
		cumulativePaths += strings.Trim(v, " ")
		if currPath != nil && currPath.Folder[v] != nil {
			currPath = currPath.Folder[v]
		} else if currPath == nil {
			rootDir = newFolder(v, cumulativePaths)
			currPath = rootDir
		} else {
			folder := newFolder(v, cumulativePaths)
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
		//fmt.Println(k)
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

func directoryExist(path string) *Folder {
	paths := deleteEmpty(strings.SplitAfter(path, "/"))
	//fmt.Println(paths)
	//name := paths[len(paths)-1]
	if len(paths) > 1 {
		paths = paths[1:len(paths)]
	}
	currPath := rootDir
	fmt.Println(paths)
	for _, v := range paths {
		if v == "../" {
			currPath = BackMap[currPath]
		} else if currPath.Folder[v] != nil {
			currPath = currPath.Folder[v]
		} else {
			return nil
		}
	}
	return currPath
}

func checkSuffix(path string) string {
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}
	return path
}

func checkPrefix(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = checkSuffix(cwd.Path) + path
	}
	return path
}
