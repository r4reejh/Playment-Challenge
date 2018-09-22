package main

import (
	"fmt"
	"strings"
)

func processCommand(cmd string) {
	args := strings.Fields(cmd)
	switch arg := args[0]; arg {
	case "cd":
		path := checkSuffix(checkPrefix(args[1]))
		fmt.Println("ss: " + path)
		d := directoryExist(path)
		if d != nil {
			cwd = d
			fmt.Println("SUCC: DIR CHANGED " + d.Path)
		} else {
			fmt.Println("ERR: DIR NOT FOUND")
		}
	case "mkdir":
		path := args[1]
		pathSplit := strings.Split(path, "/")
		name := pathSplit[len(pathSplit)-1]
		path = checkSuffix(checkPrefix(path))
		d := directoryExist(args[1])
		if d != nil {
			fmt.Println("ERROR: DIR EXISTS")
		} else {
			fmt.Println(path)
			addFolder(path, name, newFolder(checkSuffix(name), path))
			fmt.Println("SUCC: DIR ADDED")
		}
	case "ls":
		fmt.Println("DIRS:")
		printMap(cwd, 0)
	case "pwd":
		//fmt.Println("print working directory")
		fmt.Print("PWD: ")
		fmt.Println(cwd.Path)
	case "session":
		initFS()
		fmt.Println("SUCC: SESSION CLEARED")
	default:
		fmt.Println("ERR: CANNOT RECOGNIZE INPUT")
	}
}
