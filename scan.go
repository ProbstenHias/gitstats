package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

// scan the current directory for a .git folder, if there is none, move up one directory and scan there
func scan(path string) (string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() && file.Name() == ".git" {
			return path, nil
		}
	}
	if path == "/" {
		return "", errors.New("no .git folder found")
	}
	// get the parent directory of path
	parent := filepath.Dir(path)
	return scan(parent)
}
