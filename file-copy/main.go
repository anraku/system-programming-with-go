package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("ファイル名を指定してください")
		os.Exit(1)
	}

	src := args[0]
	dist := args[1]
	oldFile, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	newFile, err := os.Create(dist)
	if err != nil {
		panic(err)
	}
	io.Copy(newFile, oldFile)
}
