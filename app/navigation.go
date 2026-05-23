package main

import (
	"fmt"
	"os"
	"path/filepath"
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
	fullpath := filepath.Join(nav.wd, destination)
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
