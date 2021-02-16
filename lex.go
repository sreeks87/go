package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the morganAndString function below.
func morganAndString(a string, b string) string {
	a += "z"
	b += "z"

	aSlice := strings.Fields(a)
	bSlice := strings.Fields(b)
	newString := ""
	count := len(aSlice) + len(bSlice)
	for i := 0; i < count; i++ {
		if a > b {
			aSlice = strings.Fields(a)
			bSlice = strings.Fields(b)
			if len(aSlice) > 0 {
				newString += aSlice[0]
			} else {
				newString += ""
			}
			aSlice = aSlice[1:]
		} else {
			if len(bSlice) > 0 {
				newString += bSlice[0]
			} else {
				newString += ""
			}
			bSlice = bSlice[1:]
		}
	}
	return strings.TrimRight(newString, "z")
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		a := readLine(reader)

		b := readLine(reader)

		result := morganAndString(a, b)

		fmt.Fprintf(writer, "%s\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
