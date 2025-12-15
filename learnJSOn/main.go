package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/2"
	//oneway
	// resonse, err := http.Get(url)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer resonse.Body.Close()

	// if resonse.StatusCode == http.StatusOK {
	// 	bodyBytes, err := io.ReadAll(resonse.Body)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	todoItem := Todo{}
	// 	json.Unmarshal(bodyBytes, &todoItem)

	// 	fmt.Printf("Data from API: %+v", todoItem)
	// }

	// otherway
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode == http.StatusOK {

		todoItem := Todo{}
		decoder := json.NewDecoder(response.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&todoItem); err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("Data from API: %+v", todoItem)

		// byteData, err := json.Marshal(&todoItem)
		byteData, err := json.MarshalIndent(&todoItem, "", "\t")

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(byteData))

	}

}
