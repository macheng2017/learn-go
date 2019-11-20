package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		i int
		b bool
		s string
	)
	r := strings.NewReader("5 truegophers")
	n, err := fmt.Fscanf(r, "%d %s ", &i, &s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fscanf: %v\n", err)
	}
	fmt.Println(i, s)
	fmt.Println(n)
}
