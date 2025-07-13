package main

type JsonData struct {
	Count int `json: "int"`
	Next string `json: "string"`
	Previous string `json: "string"`
	Results []location
}

type location struct {
	Name string `json: "string"`
	Url string `json: "string"`
}

