package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type BeefResponse struct {
	Beef map[string]int `json:"beef"`
}

func getMeats() (*BeefResponse, error) {
	res, err := getTypeOfMeat()
	if err != nil {
		log.Fatalf("Failed to retrieve meats data: %v", err)
		return nil, err
	}
	var beefResponse BeefResponse
	beefResponse.Beef = map[string]int{}
	tokenized(res, &beefResponse)
	fmt.Print(beefResponse)
	return &beefResponse, nil
}

func getTypeOfMeat() (string, error) {
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	result, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	stringResult := string(result)
	return stringResult, nil
}

func tokenized(txt string, response *BeefResponse) {
	txt = strings.ReplaceAll(txt, ",", "")
	txt = strings.ReplaceAll(txt, ".", "")
	txt = strings.ReplaceAll(txt, "\n", "")
	splitedTxt := strings.Split(txt, " ")
	for _, typeMeat := range splitedTxt {
		if typeMeat != "" {
			response.Beef[typeMeat]++
		}
	}
}
