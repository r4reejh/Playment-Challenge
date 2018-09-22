package main

import (
	"log"
	"os"
	"path/filepath"
)

// function for loading the filesystem from the directory specified in initDir
func initFS() {

	initDir = "/home/"
	rootDir = nil
	BackMap = make(map[*Folder]*Folder)

	// Utility function to add the initial root directory to tree
	addFolderForce(initDir)

	// Walking the directories to load them into a tree
	err := filepath.Walk(initDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Print(err)
				return nil
			}
			if info.IsDir() {
				folder := newFolder(info.Name()+"/", path)
				name := info.Name() + "/"
				addFolder(path, name, folder)
			}
			return nil
		})

	if err != nil {
		log.Fatal(err)
	}
	cwd = rootDir
	return
}

// function to initialize the commands
func initCommands() {

	CommandMap = make(map[string]*Command)

	//creating Command variables on runtime, new additions can be easily made here
	createCommands()

}
