package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("messages.txt")

	if err != nil {
		log.Fatal("error opening file", err)
	}

	defer f.Close()

	str := ""
	for {
		data := make([]byte, 8)
		n, err := f.Read(data)

		if err != nil {
			log.Fatal("error reading file", err)
			break
		}

		data = data[:n]
		if i := bytes.IndexByte(data, '\n'); i != -1 {
			str += string(data[:i])
			data = data[i+1 : n]
			fmt.Printf("read %s\n", str)
			str = ""
		}
		str += string(data)
	}

	if len(str) != 0 {
		fmt.Printf("read %s\n", str)
	}
}
