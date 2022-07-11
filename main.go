package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/")
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var f mypost
	_ = json.Unmarshal(body, &f)
	fmt.Printf("%+v",f[0].Body)

}

type mypost []struct {
	UserID string `json:"user_id"`
	ID string `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}

