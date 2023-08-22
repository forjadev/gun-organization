package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"github.com/forjadev/gun-organization/config"
	"github.com/forjadev/gun-organization/schemas"
)

type Member struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
	URL   string `json:"url"`
}

type MemberInfo struct{}

var (
	TOKEN string
	URL   string
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	// GITHUB
	TOKEN = os.Getenv("GITHUB_TOKEN")
	URL = os.Getenv("GITHUB_URL")

	members, err := getMembers()
	if err != nil {
		panic(err)
	}

	for _, member := range members {
		err := insertOrUpdateMember(schemas.Member{GithubID: strconv.Itoa(member.Id), Login: member.Login, URL: member.URL})
		if err != nil {
			panic(err)
		}
	}

	// // FORJAA_DB

	// // Initialize configs
	// err = config.Init()
	// if err != nil {
	// 	panic(err)
	// }

	// db := config.GetDatabase()

	// var members []schemas.Member
	// db.Find(&members)

	// fmt.Println("FIND MEMBERS")

	// for _, member := range members {
	// 	fmt.Println(member)
	// }
}

func getMembers() ([]Member, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	// GITHUB
	TOKEN := os.Getenv("GITHUB_TOKEN")
	URL := os.Getenv("GITHUB_URL")

	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	var data []Member
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func insertOrUpdateMember(member schemas.Member) error {
	err := config.Init()
	if err != nil {
		return err
	}

	db := config.GetDatabase()
	result := db.Create(&member)

	if result.Error != nil {
		return result.Error
	}

	fmt.Println(result.RowsAffected)
	return nil
}
