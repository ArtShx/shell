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

func GetNavigation() *Navigation {
	once.Do(func() {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Printf("Failed getting cwd: %s\n", err)
		}
		nav = &Navigation{wd: cwd}
	})
	return nav
}

func (nav *Navigation) ChangeDirectory(destination string) {
	// os.Stat(destination)
	// filepath.exi
	//filepath.Abs(destination)
	if stat, err := os.Stat(destination); err == nil && stat.IsDir() {

		absPath, err := filepath.Abs(destination)
		// fmt.Printf("Destination: %s\n2. Abs: %s\n", destination, absPath)
		if err == nil {
			nav.wd = absPath
			os.Setenv("PWD", destination)
		} else {
			// fmt.Printf("%s\n", err)
		}
	}
}
