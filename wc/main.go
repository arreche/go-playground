package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func countWords(s string) int {
	return len(strings.Fields(s))
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var b []byte
	var err error
	if len(os.Args) > 1 {
		b, err = ioutil.ReadFile(os.Args[1])
	} else {
		b, err = ioutil.ReadAll(os.Stdin)
	}
	checkErr(err)
	count := countWords(string(b))
	fmt.Print(count)
}
