package main

import "fmt"

func main() {
	fmt.Println("Maps")

	language := make(map[string]string)

	language["JS"] = "JavaScript"
	language["RB"] = "Ruby"
	language["PY"] = "Python"

	fmt.Println("List of Langs , ", language)
	fmt.Println("JS :", language["JS"])

	delete(language, "RB")
	fmt.Println("List of Langs , ", language)

	//loops
	for key, value := range language {
		fmt.Printf("For key %v , Value is %v ", key, value)
	}
	for _, value := range language {
		fmt.Printf("Value is %v ", value)
	}
}
