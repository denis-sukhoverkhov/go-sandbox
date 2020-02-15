package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	user   string = "tt"
	home   string = "rr"
	gopath string = "ee"
)

func init() {
	if user == "" {
		log.Fatal("$USER not set")
	}
	if home == "" {
		home = "/home/" + user
	}
	if gopath == "" {
		gopath = home + "/go"
	}
	// gopath may be overridden by --gopath flag on command line.
	flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
}

func main() {
	fmt.Println("test")
}
