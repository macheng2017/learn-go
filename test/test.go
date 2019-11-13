package main

import (
	"fmt"
	"os"
)

func main() {
	const name, age = "Kim", 22
	fmt.Fprintln(os.Stdout, name, "is", age, "years old.")

	// The n and err return values from Fprintln are
	// those returned by the underlying io.Writer.
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Fprintln: %v\n", err)
	//}
	//fmt.Println(n, "bytes written.")
}
