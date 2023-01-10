package main

import (
	"log"
	"os"
	"os/user"

	"github.com/BurntSushi/toml"
)

// Info from config file
type Config struct {
	Email string
}

// getDotFilePath returns the dot file for the repos list.
// Creates it and the enclosing folder if it does not exist.
func getDotFilePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	dotFile := usr.HomeDir + "/.gitstats"

	_, err = os.Stat(dotFile)
	if err == nil {
		return dotFile
	}
	f, err := os.Create(dotFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	return dotFile
}

// Reads info from config file
func ReadConfig() Config {
	var configfile = getDotFilePath()
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	return config
}
