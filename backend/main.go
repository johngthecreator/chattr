package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// struct for the data type I'm getting back
type ApiResponse struct {
	Results []Quote `json:"results"`
}

type Quote struct {
	Phrase string `json:"content"`
	Author string `json:"author"`
}

func main() {
	resp, err := http.Get("https://api.quotable.io/search/quotes?query=hope")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// this initializes the variable that the unmarshalled json data will be returned to
	var data ApiResponse

	// this throws an error if the error isn't nil and the err is being assigned to the json.unmarshal function
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(data.Results[0].Phrase)
}
