package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("hehe\n")
	os.Mkdir("robodep", 0700)
}
