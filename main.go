package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	var email = flag.String("email", "your@email.com", "the email to scan for")
	flag.Parse()
	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	if flag.NArg() > 0 {
		if flag.Arg(0)[0] != '/' && flag.Arg(0)[0] != '~' {
			mydir = mydir + "/" + flag.Arg(0)
		} else {
			mydir = flag.Arg(0)
		}

	}
	fmt.Println(flag.NArg())
	fmt.Println(mydir)
	fmt.Println(*email)
	directory, err := scan(mydir)
	if err != nil {
		log.Fatalf(
			"no .git folder found, please run this command from a directory with a .git folder",
		)
	}
	stats(*email, directory)

}
