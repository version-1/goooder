package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) <= 0 {
		useage()
		return
	}

	cmd(args[0])
}

func useage() {
	fmt.Println("goooder init")
}

func cmd(command string) {
	switch command {
	case "init":
	case "gen":
	default:
		useage()
	}

}
