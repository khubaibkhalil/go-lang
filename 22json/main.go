package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// EncodeJson()
	DecodeJson()
}

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platfrom string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func EncodeJson() {

	courses := []course{
		{"ReactJs", 299, "yt.com", "abc123", []string{"web", "Js"}},
		{"Go-Lang", 400, "yt.com", "abc123", []string{"backend", "go"}},
		{"MERN", 400, "yt.com", "abcd123", nil},
	}

	// finalJson, err := json.Marshal(courses)
	finalJson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
        {
                "coursename": "ReactJs",
                "price": 299,
                "website": "yt.com",
                "tags": ["web","Js"]
        }
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("Json is Valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v", lcoCourse)
	} else {
		fmt.Println("Invlid Json")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("\n%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and Type is: %T \n", k, v, v)
	}

}
