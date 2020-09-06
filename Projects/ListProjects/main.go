package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	url := os.Args[1]
	apiKey := os.Args[2]

	response, err := http.Get(url + "/api/projects")
	response.Header.Set("X-Octopus-ApiKey", apiKey)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	output, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", output)

}
