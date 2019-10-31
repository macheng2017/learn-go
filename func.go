package learngo

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// go 函数的参数列表和返回值都是类型在后
func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported operation: " + op)
	}
}

// 双返回值的应用场景,第二返回值返回错误信息
func eval2(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
		//panic("unsupported operation: " + op)
	}
}

// go函数可以返回两个值
func div(a, b int) (int, int) {
	return a / b, a % b
}

func div2(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+"(%d, %d)", opName, a, b)
	return op(a, b)
} //Calling function main.pow with args (3, 4)81
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
func main() {

	//fmt.Println(eval(3, 4, "ii"))
	q, r := div2(13, 8)
	fmt.Println(q, r)

	if res, err := eval2(3, 4, "%"); err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println(res)
	} //Error unsupported operation: %

	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))
	// Calling function main.main.func1 with args (3, 4)81
}
