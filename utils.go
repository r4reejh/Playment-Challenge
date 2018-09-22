package main

import (
	"strings"
)

// utility function to create new instances of Folder
func newFolder(name string, path string) *Folder {
	return &Folder{name, make(map[string]*Folder), path}
}

// adding folder from absolute path
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

// force adding folder, similar to `mkdir -p` command
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

// deleting empty values from slice
func deleteEmpty(s []string) []string {

	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r

}

// check if directory exists for absolute path, return address if exists
func directoryExist(path string) *Folder {

	paths := deleteEmpty(strings.SplitAfter(path, "/"))
	if len(paths) > 1 {
		paths = paths[1:len(paths)]
	}
	currPath := rootDir
	for _, v := range paths {
		if v == "/" {
			return currPath
		} else if v == "../" {
			currPath = BackMap[currPath]
		} else if currPath.Folder[v] != nil {
			currPath = currPath.Folder[v]
		} else {
			return nil
		}
	}
	return currPath

}

// appending trailing `/`
func checkSuffix(path string) string {

	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}
	return path
}

// prepending current working directory in case of relative path as input to generate absilute path
func checkPrefix(path string) string {

	if !strings.HasPrefix(path, "/") {
		path = checkSuffix(cwd.Path) + path
	}
	return path
}

// remove suffix
func removeSuffix(path string) string {
	if len(path) > 1 && strings.HasSuffix(path, "/") {
		path = path[0 : len(path)-1]
	}
	return path
}

// remove prefix
func removePrefix(path string) string {
	if len(path) > 1 && strings.HasPrefix(path, "/") {
		path = path[1:len(path)]
	}
	return path
}
