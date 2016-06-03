package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"log"
	"flag"
)

type Row struct {
	Text []string
}
func main() {
	file := flag.String("path", "", "file path")
	padding := flag.String("padding", "", "print padding")

	flag.Parse()

	if (*file == "") {
		fmt.Println("file", *file, "not exist")

		os.Exit(1)
	}

	if (*padding == "") {
		*padding = "-20";
	}


	csv := parse(*file)

	fmt.Println("")
	for _,element := range csv {

		for _,elem := range element.Text {
			fmt.Printf("%" + *padding + "s", elem)
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func parse(file string) []Row {
	var result []Row

	reader, err := os.Open(file)
	if err != nil {
		log.Panic(err)
	}

	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			temp := strings.Split(line, ";")
			result = append(result, Row{temp})
		}
	}

	return result
}
