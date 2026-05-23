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
	if stat, err := os.Stat(fullpath); err == nil && stat.IsDir() {
		absPath, err := filepath.Abs(fullpath)
		if err != nil {
			fmt.Printf("cd: %s: No such file or directory\n", fullpath)
			return
		}
		os.Chdir(absPath)
		nav.wd = absPath
		os.Setenv("PWD", fullpath)
	} else {
		fmt.Printf("cd: %s: No such file or directory\n", destination)
	}
}
