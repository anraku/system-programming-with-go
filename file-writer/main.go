package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	f, err := os.OpenFile("/Users/daimori/play/go/system-programming/file-writer/memo.txt", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	now := time.Now()
	memo := "今日は楽しかった"
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(f, "%s %s\n", now.Format("2006-01-02 15:04:05"), memo)
}
