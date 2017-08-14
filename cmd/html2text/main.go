package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/TheBookPeople/html2text"
)

func main() {
	info, _ := os.Stdin.Stat()

	if info.Mode()&os.ModeNamedPipe == 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage:")
		fmt.Println("  cat yourfile.html | html2text")
		os.Exit(-1)
	}

	data, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println(html2text.HTML2TextCustomLine(string(data), "\r\n"))

}
