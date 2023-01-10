package main

import (
	"flag"
	"os"
)

func myFlagParse() {
	arangeArgs()
	flag.Parse()
}

func arangeArgs() {
	flags := make([]string, 0, len(os.Args))
	args := make([]string, 0, len(os.Args))
	for _, arg := range os.Args[1:] {
		if arg[0] == '-' {
			flags = append(flags, arg)
		} else {
			args = append(args, arg)
		}
	}
	for i, flag := range flags {
		os.Args[i+1] = flag
	}
	argsOffset := len(flags) + 1
	for i, arg := range args {
		os.Args[i+argsOffset] = arg
	}

}
