package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("input1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

    for {
        line, _, err := reader.ReadLine();

        if err == io.EOF {
            break
        }
        fmt.Println(line)
    }

}
