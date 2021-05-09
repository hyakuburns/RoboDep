package main

import (
	"os"
)

func main() {
	if !FileExistence("robodep") {
		os.Mkdir("robodep", 0700)
	}
	var floc string = "./dep.robo"

	argv := os.Args
	argc := len(os.Args)
	if argc <= 1 {
		Instructions()
		os.Exit(1)
	}
	if argc < 3 && argv[1] != "up" {
		Instructions()
		os.Exit(1)

	}

	switch argv[1] {
	case "git":
		GitAddRepo(floc, argv)
	case "hg":
		HgAddRepo(floc, argv)
	case "up":
		ParseDeps(floc)
	default:
		Instructions()
	}
}
