package main

import (
	"fmt"
	"os"
	"sync"
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
		fmt.Printf("%s\n", nav.wd)
	})
	return nav
}
