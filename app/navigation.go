package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	// "path/filepath"
)

type Navigation struct {
	wd string
}

var (
	nav  *Navigation
	once sync.Once
)

func _getWd() string {
	cwd := os.Getenv("PWD")
	return cwd
}

func GetNavigation() *Navigation {
	once.Do(func() {
		cwd := _getWd()
		nav = &Navigation{wd: cwd}
	})
	nav.wd = _getWd()
	return nav
}

func ChangeDirectory(destination string) {
	nav := GetNavigation()
	var fullpath string
	if strings.HasPrefix(destination, "/") {
		fullpath = destination
	} else if strings.HasPrefix(destination, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Failed to get home dir: %s\n", err)
			return
		}
		fullpath = strings.Replace(destination, "~", home, 1)
	} else {
		fullpath = filepath.Join(nav.wd, destination)
	}
	absPath, err := filepath.Abs(fullpath)
	if err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", fullpath)
		return
	}
	err = os.Chdir(absPath)
	if err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", destination)
		return
	}
	nav.wd = absPath
	os.Setenv("PWD", fullpath)
}
