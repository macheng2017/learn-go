package learngo

import "fmt"

func main() {

	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	m2 := make(map[string]string) // m2 === empty map
	var m3 map[string]string      // m3 === nil
	fmt.Println(m, m2, m3)
	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}
	// map 是一个无序的底层实现是hashmap
	fmt.Println("Getting values")
	// 这里course 取到的就是v ,而ok取到的是函数返回的第二个返回值是一个布尔值

	courseName, ok := m["course"]
	fmt.Println(courseName, ok)

	if courseName, ok := m["cause"]; ok {
		fmt.Println(courseName)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("deleting ")

	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)

}
