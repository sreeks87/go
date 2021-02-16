package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the beautifulDays function below.
func beautifulDays(i int32, j int32, k int32) int32 {
	var c int32 = 0
	for n := i; n <= j; n++ {
		if (abs(n-rev(n)) % k) == 0 {
			c++
		}
	}
	return c

}
func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}
func rev(n int32) int32 {
	var newnum int32 = 0
	for {
		rem := n % 10
		newnum = newnum*10 + rem
		n = n / 10
		if n == 0 {
			break
		}
	}
	return newnum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	ijk := strings.Split(readLine(reader), " ")

	iTemp, err := strconv.ParseInt(ijk[0], 10, 64)
	checkError(err)
	i := int32(iTemp)

	jTemp, err := strconv.ParseInt(ijk[1], 10, 64)
	checkError(err)
	j := int32(jTemp)

	kTemp, err := strconv.ParseInt(ijk[2], 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := beautifulDays(i, j, k)

	fmt.Fprintf(writer, "%d\n", result)

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
