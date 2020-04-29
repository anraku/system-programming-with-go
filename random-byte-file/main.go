package main

import (
	"crypto/rand"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("rand.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := make([]byte, 1024)
	n, _ := rand.Read(buf)
	fmt.Println(n)
	f.Write(buf)
}
