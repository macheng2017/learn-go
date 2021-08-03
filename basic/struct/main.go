package main

import "fmt"

func main() {
	cases := []struct {
		name string
		age  int32
		op   func() (string, error)
	}{
		{"v",
			32,
			func() (string, error) {
				return "fun1", nil
			},
		},
		{
			"mailv2",
			30,
			func() (string, error) {
				return "fun2", nil
			},
		},
	}
	//for i := range cases {
	//
	//}

	fmt.Printf("%+v", cases)
	for _, s := range cases {
		str, _ := s.op()

		fmt.Printf("%+v \n", str)
	}

	fmt.Printf("%+v \n", cases)
}
