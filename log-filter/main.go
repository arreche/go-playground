package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	path := flag.String("path", "myapp.log", "The path to the log file to scan")
	level := flag.String("level", "ERROR", "Log level to filter (INFO | WARNING | ERROR)")
	flag.Parse()
	f, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		l, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(l, *level) {
			fmt.Println(l)
		}
	}
}
