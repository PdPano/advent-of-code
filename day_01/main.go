package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
    part := 1;
    if (len(os.Args) > 1 && os.Args[1]=="2"){
        part = 2
    }

	file, err := os.Open("input1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
    acc := 0;

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

        if(part == 2){
            line = convert_number_names(line)
        }

        r := keep_numbers(line)
        final, err := strconv.Atoi(string([]byte{r[0], r[len(r)-1]}));
        if err != nil {
            panic(err)
        }
        acc += final
	}
    fmt.Println(acc)

}

func keep_numbers(inp []byte) []byte {
	output := []byte{}

	for _, v := range inp {
		if v >= '0' && v <= '9' {
            output = append(output, v)
		}
	}
    return output
}

func convert_number_names(inp []byte) []byte {
    conv := string(inp)
    conv = strings.ReplaceAll(conv, "zero", "zero0zero");
    conv = strings.ReplaceAll(conv, "one", "one1one");
    conv = strings.ReplaceAll(conv, "two", "two2two");
    conv = strings.ReplaceAll(conv, "three", "three3three");
    conv = strings.ReplaceAll(conv, "four", "four4four");
    conv = strings.ReplaceAll(conv, "five", "five5five");
    conv = strings.ReplaceAll(conv, "six", "six6six");
    conv = strings.ReplaceAll(conv, "seven", "seven7seven");
    conv = strings.ReplaceAll(conv, "eight", "eight8eight");
    conv = strings.ReplaceAll(conv, "nine", "nine9nine");
    return []byte(conv)
}
