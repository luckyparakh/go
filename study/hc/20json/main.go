package main

import (
	"encoding/json"
	"fmt"
)

// //Plain
// type course struct {
// 	Name     string
// 	Price    int
// 	Platform string
// 	Password string
// 	Tags     []string
// }

type course struct {
	Name     string `json:"courses"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("JSON")
	jsonEncode()
	deCodeJson()
}

func jsonEncode() {
	courses := []course{
		{"React JS", 299, "lco.in", "abc123", []string{"web", "js"}},
		{"MERN JS", 199, "lco.in", "ddd123", []string{"full stack", "js"}},
		{"Angular JS", 199, "lco.in", "bbb123", nil},
	}
	//package this in jsom

	finalJson, err := json.Marshal(courses)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", finalJson)

	finalJsonM, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", finalJsonM)
}

func deCodeJson() {
	jsonData := []byte(`
		{
			"courses": "React JS",
			"Price": 299,
			"website": "lco.in",
			"tags": ["web","js"]
		}
	`)

	var lcoC course
	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("Json Decode")
		json.Unmarshal(jsonData, &lcoC)
		fmt.Printf("%#v\n", lcoC)
	} else {
		fmt.Println("Invalid Json")
	}
	// User interface for values because value can be of any data type like it can be slice, float etc
	var myData map[string]interface{}

	json.Unmarshal(jsonData, &myData)
	fmt.Println(myData)

}
