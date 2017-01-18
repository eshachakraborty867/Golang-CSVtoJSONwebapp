package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// create a http client
	// fetch contents of web page
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://www.google.com", nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error resolving host")
		return
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.Status)
	fmt.Println(string(respBody))

	// parse html
}
