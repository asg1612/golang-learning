package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, err := http.Get("http://google.com")
	if err != nil {
		fmt.Printf("The http resquest failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
