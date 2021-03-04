package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getPost() []byte {
	fmt.Println("Fetching data from remote...")

	resp, err := http.Get(os.Getenv("API_HOST"))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Got data from NASA")

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return respBody
}
