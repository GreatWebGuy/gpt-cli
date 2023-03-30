package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	apiKey          = "your_api_key_here"
	endpoint = "https://api.openai.com/v1/engines/davinci/completions"
	contentTypeJSON = "application/json"
)

type Payload struct {
	Prompt      string `json:"prompt"`
	MaxTokens   int    `json:"max_tokens"`
	Temperature float64 `json:"temperature"`
}

type ResponseData struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gpt [prompt]")
		os.Exit(1)
	}

	prompt := strings.Join(os.Args[1:], " ")

	payload := Payload{
		Prompt:      prompt,
		MaxTokens:   50,
		Temperature: 0.7,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Content-Type", contentTypeJSON)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	var responseData ResponseData
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if len(responseData.Choices) > 0 {
		fmt.Print(strings.TrimSpace(responseData.Choices[0].Text))
	} else {
		fmt.Println("No response received")
	}
}
