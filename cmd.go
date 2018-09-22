package main

import (
	"fmt"
	"strings"
)

func processCommand(cmd string) {
	args := strings.Fields(cmd)
	switch arg := args[0]; arg {
	case "cd":
		if len(args) > 1 {
			path := checkSuffix(checkPrefix(args[1]))
			d := directoryExist(path)
			if d != nil {
				cwd = d
				fmt.Println("SUCC: REACHED")
			} else {
				fmt.Println("ERR: INVALID PATH")
			}
		}
	case "mkdir":
		for i := 1; i < len(args); i++ {
			path := args[i]
			pathSplit := strings.Split(path, "/")
			name := pathSplit[len(pathSplit)-1]
			path = checkSuffix(checkPrefix(path))
			d := directoryExist(path)
			if d != nil {
				fmt.Println("ERR: DIRECTORY ALREADY EXISTS")
				return
			}
			addFolder(path, name, newFolder(checkSuffix(name), path))
		}
		fmt.Println("SUCC: CREATED")
	case "ls":
		fmt.Print("DIRS: ")
		//printMap(cwd, 0)
		printLS(cwd)
	case "pwd":
		//fmt.Println("print working directory")
		fmt.Print("PATH: ")
		fmt.Println(cwd.Path)
	case "rm":
		if len(args) > 1 {
			path := checkSuffix(checkPrefix(args[1]))
			d := directoryExist(path)
			if d != nil {
				ok := deleteDirectory(d)
				if ok {
					fmt.Println("SUCC: DELETED")
				} else {
					fmt.Println("ERR")
				}
			} else {
				fmt.Println("ERR: INVALID PATH")
			}
		}
	case "session":
		initFS()
		fmt.Println("SUCC: CLEARED: RESET TO ROOT")
	default:
		fmt.Println("ERR: CANNOT RECOGNIZE INPUT")
	}
}
