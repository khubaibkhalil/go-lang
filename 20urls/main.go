package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://www.google.com:3000/learn?coursename=reactjs&paymentid=322342"

func main() {

	//parse
	result, _ := url.Parse(myurl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.RawQuery)
	// fmt.Println(result.Port())
	qparams := result.Query()

	fmt.Println(qparams)

	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])

	for _, val := range qparams {
		fmt.Println(val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "www.google.com",
		Path:     "/search",
		RawQuery: "id=1",
	}

	anOtherUrl := partsOfUrl.String()

	fmt.Println(anOtherUrl)

}
