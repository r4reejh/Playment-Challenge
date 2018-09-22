package main

import (
	"fmt"
	"log"
	"os"
)

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
