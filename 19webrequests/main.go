package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.google.com/"

func main() {
	fmt.Println("Web Request")
	res, err := http.Get(url)
	checkNilErr(err)
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	checkNilErr(err)
	content := string(data)
	fmt.Println(content)

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
