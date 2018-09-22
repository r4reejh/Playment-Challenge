package main

import "fmt"

// Folder struct type used to store Folder heirarchy in tree
type Folder struct {
	Name   string
	Folder map[string]*Folder
	Path   string
}

// BackMap struct type to store parent directory address
var BackMap map[*Folder]*Folder

//
func (folder *Folder) returnLS() string {
	out := ""
	for _, v := range folder.Folder {
		out += removePrefix(removeSuffix(v.Name)) + " "
	}
	return out
}

func (folder *Folder) printTree(spaces int) {
	for i := 0; i < spaces; i++ {
		fmt.Print(" ")
	}
	fmt.Println(folder.Name)
	for _, v := range folder.Folder {
		v.printTree(spaces + 1)
		//fmt.Println(k)
	}
}

func (folder *Folder) deleteDirectory() bool {
	parent, ok := BackMap[folder]
	if ok {
		if cwd == folder {
			cwd = parent
		}
		delete(BackMap, folder)
		delete(parent.Folder, folder.Name)
		folder = nil
		return true
	}
	return false
}
