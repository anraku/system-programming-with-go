package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		os.Exit(1)
	}
	n, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}

	exit := make(chan int)
	go func() {
		<-time.After(time.Duration(n) * time.Second)
		close(exit)
	}()
	<-exit
	fmt.Println("finish")
}
