package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

var keywords = [7]string{"If", "For", "While", "Func", "Class", "Extends", "Singleton"}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("./test.dabu")
	check(err)
	scriptReader := strings.NewReader(string(dat))
	for {
		c, err := scriptReader.ReadByte()
		if err == io.EOF {
			break
		} else {
			fmt.Print(string(c))
		}
	}
}
