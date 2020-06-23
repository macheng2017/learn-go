package infra

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Retrieve struct {
}

func (Retrieve) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	return string(bytes)
}
