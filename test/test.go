package main

import "fmt"

func main() {
	//var (
	//	i int
	//	b bool
	//	s string
	//)
	//r := strings.NewReader("5 truegophers")
	//n, err := fmt.Fscanf(r, "%d %s ", &i, &s)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Fscanf: %v\n", err)
	//}
	//fmt.Println(i, s)
	//fmt.Println(n)

	for i := 1; i <= 8; i++ {
		b := 0
		n1 := b + 17*i
		n2 := 246 - 15*i
		n3 := 156 - 13*i

		fmt.Printf("num1: %d  ", n1)
		fmt.Printf("num2: %d  ", n2)
		fmt.Printf("num3: %d \n ", n3)
		//fmt.Printf("num3: %d \n ", 0+13*i)
	}

}
