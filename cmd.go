package main

import (
	"fmt"
	"strings"
)

func processCommand(cmd string) {
	args := strings.Fields(cmd)
	//fmt.Println(args)
	switch arg := args[0]; arg {
	case "cd":
		d := directoryExist(args[1])
		if d != nil {
			cwd = d
			fmt.Println("SUCC: DIR CHANGED")
		}
	case "mkdir":
		path := args[1]
		pathSplit := strings.Split(path, "/")
		name := pathSplit[len(pathSplit)-1]
		if !strings.HasPrefix(path, "/") {
			path = cwd.Path + path
		}
		d := directoryExist(args[1])
		if d != nil {
			fmt.Println("ERROR: DIR EXISTS")
		} else {
			addFolder(path, name, cwd)
			fmt.Println("SUCC: DIR ADDED")
		}
	case "ls":
		fmt.Println("show files and folders")
		printMap(rootDir, 0)
	case "pwd":
		//fmt.Println("print working directory")
		fmt.Println(cwd.Path)
	case "session":
		fmt.Println("clear current session")
	default:
		fmt.Println("ERR: CANNOT RECOGNIZE INPUT")
	}
}
