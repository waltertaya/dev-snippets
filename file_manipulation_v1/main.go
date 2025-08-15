package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("Error opening the messages.txt:", err)
	}
	defer f.Close()

	lines := getLinesChannel(f)

	for line := range lines {
		fmt.Println("read:", line)
	}
}

func getLinesChannel(f io.Reader) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)

		var str string

		for {
			data := make([]byte, 8)
			n, err := f.Read(data)

			if n > 0 {
				data = data[:n]
				for {
					i := bytes.IndexByte(data, '\n')
					if i == -1 {
						str += string(data)
						break
					}

					str += string(data[:i])
					ch <- str

					str = ""

					data = data[i+1:]
				}
			}

			if err == io.EOF {
				if len(str) > 0 {
					ch <- str
				}
				break
			}

			if err != nil {
				log.Println("Read error:", err)
				break
			}
		}
	}()

	return ch
}
