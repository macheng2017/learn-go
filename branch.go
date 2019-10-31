package learngo

import (
	"fmt"
	"io/ioutil"
	"log"
)

func branch() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

}
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score <= 60:
		g = "D"
	case score <= 80:
		g = "C"
	case score <= 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	//branch()
	fmt.Println(
		//grade(-2),
		grade(50),
		grade(60),
		grade(70),
		grade(80),
		grade(100),
	)
}
