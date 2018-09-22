package main

import (
	"fmt"
	"strings"
)

// Command to store the different functions and their details
type Command struct {
	key    string
	action func([]string) string
}

// CommandMap to Map the commands available
var CommandMap map[string]*Command

// function used to create the Commands and store them in the Map at runtime
func newCommand(key string, action func([]string) string) bool {
	CommandMap[key] = &Command{key, action}
	_, ok := CommandMap[key]
	if ok {
		return true
	}
	return false
}

// process the user input
func processCommand(cmd string) {
	args := strings.Fields(cmd)
	f, ok := CommandMap[args[0]]
	if ok {
		res := f.action(args)
		if res != "" {
			fmt.Println(res)
		}
	} else {
		fmt.Println("ERR: CANNOT RECOGNIZE INPUT")
	}
}

// function used in tests
func processCommandTest(cmd string) string {
	args := strings.Fields(cmd)
	f, ok := CommandMap[args[0]]
	if ok {
		res := f.action(args)
		if res != "" {
			return res
		}
		return ""
	}
	return "ERR: CANNOT RECOGNIZE INPUT"
}

// function to create the predefined commands, more commands can be created easily later
func createCommands() {

	newCommand("cd", func(args []string) string {
		if len(args) > 1 {
			path := checkSuffix(checkPrefix(args[1]))
			d := directoryExist(path)
			if d != nil {
				cwd = d
				return "SUCC: REACHED"
			}
			return "ERR: INVALID PATH"

		}
		return ""
	})

	newCommand("mkdir", func(args []string) string {
		for i := 1; i < len(args); i++ {
			path := args[i]
			pathSplit := strings.Split(path, "/")
			name := pathSplit[len(pathSplit)-1]
			path = checkSuffix(checkPrefix(path))
			d := directoryExist(path)
			if d != nil {
				return "ERR: DIRECTORY ALREADY EXISTS"
			}
			addFolder(path, name, newFolder(checkSuffix(name), path))
		}
		if len(args) > 1 {
			return "SUCC: CREATED"
		}
		return ""
	})

	newCommand("ls", func(args []string) string {
		return "DIRS: " + cwd.returnLS()
	})

	newCommand("pwd", func(args []string) string {
		return "PATH: " + cwd.Path
	})

	newCommand("rm", func(args []string) string {
		for i := 1; i < len(args); i++ {
			path := checkSuffix(checkPrefix(args[1]))
			d := directoryExist(path)
			if d != nil {
				ok := d.deleteDirectory()
				if !ok {
					return "ERR"
				}
			} else {
				return "ERR: INVALID PATH"
			}
		}
		if len(args) > 1 {
			return "SUCC: DELETED"
		}
		return ""
	})

	newCommand("session", func(args []string) string {
		initFS()
		return "SUCC: CLEARED: RESET TO ROOT"
	})
}
