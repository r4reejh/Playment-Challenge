package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func initFS() {
	rootDir = nil
	BackMap = make(map[*Folder]*Folder)
	addFolderForce(initDir)
	err := filepath.Walk(initDir, loadFileSystem)
	if err != nil {
		log.Fatal(err)
	}
	cwd = rootDir
	return
}

func loadFileSystem(path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Print(err)
		return nil
	}
	if info.IsDir() {
		currDir = path
		folder := &Folder{info.Name() + "/", make(map[string]*Folder), path}
		name := info.Name() + "/"
		addFolder(path, name, folder)
	}
	return nil
}

func indentPrint(level int) {
	for i := 0; i < level; i++ {
		fmt.Print(" ")
	}
}
