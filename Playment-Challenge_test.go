package main

import (
	"strings"
	"testing"
)

type testTable struct {
	cmd string
	out string
}

func TestOutput(t *testing.T) {
	initFS()
	initCommands()
	vs := []testTable{
		{"mkdir abcd", "SUCC: CREATED"},
		{"mkdir abcd/home", "SUCC: CREATED"},
		{"cd abcd", "SUCC: REACHED"},
		{"pwd", "PATH: /abcd"},
		{"ls", "DIRS: home"},
		{"mkdir hello world", "SUCC: CREATED"},
		{"rmdir hello", "ERR: CANNOT RECOGNIZE INPUT."},
		{"rm hello world", "SUCC: DELETED"},
		{"cd", ""},
		{"mkdir test", "SUCC: CREATED"},
		{"mkdir test", "ERR: DIRECTORY ALREADY EXISTS"},
		{"cd test", "SUCC: REACHED"},
		{"rm hello", "ERR: INVALID PATH"},
		{"mkdir testing", "SUCC: CREATED"},
		{"ls", "DIRS: testing"},
		{"asdf asdf", "ERR: CANNOT RECOGNIZE INPUT."},
		{"cd /", "SUCC: REACHED"},
		{"cd abcd", "SUCC: REACHED"},
		{"session clear", "SUCC: CLEARED: RESET TO ROOT"},
		{"pwd", "PATH: /"},
	}

	for _, v := range vs {
		out := processCommandTest(v.cmd)
		if strings.Trim(out, " ") != v.out {
			t.Errorf("Incorrect Output for cmd `%s` expected:`%s` got `%s`", v.cmd, v.out, out)
		}
	}
}
