package main

import (
	"fmt"
	"log"
	"os"

	"github.com/juju/gnuflag"
)

func main() {

	var email = gnuflag.String("email", "your@email.com", "the email to scan for")
	gnuflag.Parse(true)
	if *email == "your@email.com" {
		var conf = ReadConfig()
		email = &conf.Email
	}

	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	if gnuflag.NArg() > 0 {
		if gnuflag.Arg(0)[0] != '/' && gnuflag.Arg(0)[0] != '~' {
			mydir = mydir + "/" + gnuflag.Arg(0)
		} else {
			mydir = gnuflag.Arg(0)
		}
	}
	directory, err := scan(mydir)
	if err != nil {
		log.Fatalf(
			"no .git folder found, please run this command from a directory with a .git folder",
		)
	}
	stats(*email, directory)

}
