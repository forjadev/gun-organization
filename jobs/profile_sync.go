package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type member struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
	URL   string `json:"url"`
}

func main() {
	err := godotenv.Load(".env")

	var TOKEN = os.Getenv("GITHUB_TOKEN")
	var URL = os.Getenv("GITHUB_URL")

	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", TOKEN))
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var data []member
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	for _, member := range data {
		fmt.Println(member)
	}
}
