package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON in golang!!!")

	//EncodeJson()

	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "Learngoprogramming", "abc123", []string{"reactjs", "nodejs", "express"}},
		{"MERN Bootcamp", 199, "Learngoprogramming", "av3123", []string{"full-stack", "nodejs", "express"}},
		{"Angular Bootcamp", 239, "Learngoprogramming", "abc123", nil},
	}

	//package this data as JSON Data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", finalJson)
	fmt.Printf("length of the json %d \n", len(finalJson))
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
        {
                "courseName": "ReactJS Bootcamp",
                "Price": 299,
                "website": "Learngoprogramming",
                "tags": [
                        "reactjs",
                        "nodejs",
                        "express"
                ]
        }
	`)
	var lcoCourses course
	checkJson := json.Valid(jsonDataFromWeb)
	if checkJson {
		fmt.Println("This is valid json")
		json.Unmarshal(jsonDataFromWeb, &lcoCourses)
		fmt.Printf("%#v\n", lcoCourses)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	//some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("Key is %v and value is %v & type %T \n", key, value, value)
	}
}
