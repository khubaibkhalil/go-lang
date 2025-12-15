package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web")

	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"

	res, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println("Status Code", res.StatusCode)
	fmt.Println("Content Length", res.ContentLength)

	content, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var responseString strings.Builder
	responseString.Write(content)

	// responseString.WriteString()

	fmt.Println(responseString.String())

}

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:8000/post"

	reqBody := strings.NewReader(`
	{
		"country":"PK",
		"code": 92
	}
	`)

	res, err := http.Post(myUrl, "application/json", reqBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)
	var responseString strings.Builder
	responseString.Write(content)

	fmt.Println(responseString.String())

}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstname", "khubaib")
	data.Add("lastname", "khalil")
	data.Add("email", "khubaib@gmail.com")

	res, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)

	var responseString strings.Builder

	responseString.Write(content)

	fmt.Println(responseString.String())

}
