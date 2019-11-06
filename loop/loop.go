package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func ConventToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		x := n % 2
		result = strconv.Itoa(x) + result

	}
	return result
}
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	printFileContents(file)
}

func printFileContents(file io.Reader) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(ConventToBin(13))
	printFile("abc.txt")
	s := `sdfsd
				sdfsd

sdfds

`
	printFileContents(strings.NewReader(s))

}
