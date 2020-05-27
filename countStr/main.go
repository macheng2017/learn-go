package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	//dir, err := ioutil.ReadDir("D:/")
	//if err != nil {
	//	fmt.Printf("reading file %s ", err.Error())
	//}
	//for _, info := range dir {
	//	if !info.IsDir() {
	//		fmt.Printf(" %s \n", info.Name())
	//	}
	//
	//}

	var fileList []string
	err := filepath.Walk("D:/", visit(&fileList))

	if err != nil {
		log.Println(err)
	}
	for _, s := range fileList {
		fmt.Printf("path %s  \n", s)
	}

}

// 柯里化
func visit(fileList *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 不要在for循环当中做io操作,况且是递归
		*fileList = append(*fileList, path)
		//fmt.Printf("path %s size %d \n", path, info.Size())
		return nil
	}
}
