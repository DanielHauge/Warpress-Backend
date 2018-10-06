package main

type ExampleInput struct {
	ExampleString string `json:"example_string"`
	ExampleInteger int `json:"example_integer"`
}

type ExampleOutput struct {
	ExampleListOfIntergers []int `json:"example_list_of_intergers"`
}

type CharInfo struct {
	Name string
	Realm string
	Locale string
}



type LoginInfo struct {
	Main CharInfo
	Alts []CharInfo
}

