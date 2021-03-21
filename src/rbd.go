package main

import (
	"git.sr.ht/~hyakuburns/robodep/src/pt"
	"os"
)

func main() {
	if !pt.FileExistence("robodep") {
		os.Mkdir("robodep", 0700)
	}
	var floc string = "./dep.robo"

	argv := os.Args
	argc := len(os.Args)
	if argc <= 1 {
		pt.Instructions()
		os.Exit(1)
	}
	if argc < 3 && argv[1] != "up" {
		pt.Instructions()
		os.Exit(1)

	}

	switch argv[1] {
	case "git":
		pt.GitAddRepo(floc, argv)
	case "hg":
		pt.HgAddRepo(floc, argv)
	case "up":
		pt.ParseDeps(floc)
	default:
		pt.Instructions()
	}
}
