package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"

	"github.com/forjadev/gun-organization/config"
	"github.com/forjadev/gun-organization/schemas"
)

type Member struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
	URL   string `json:"html_url"`
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

	err = config.Init()
	if err != nil {
		panic(err)
	}

	teams, err := getTeams()
	if err != nil {
		panic(err)
	}

	for _, team := range teams {
		members, err := getMembers(team.Name)
		if err != nil {
			panic(err)
		}

		for _, member := range members {
			insertOrUpdateMember(schemas.Member{GithubID: strconv.Itoa(member.Id), Login: member.Login, URL: member.URL, TeamID: int(team.ID)})
		}
	}

	// Apenas para debug(saber oq ta gravando)
	getDbMembers()
}

func getMembers(teamName string) ([]Member, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/members", URL, teamName), nil)
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
		panic(err)
	}
	var data []Member
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	return data, nil
}

func insertOrUpdateMember(member schemas.Member) {
	db := config.GetDatabase()
	if db.Model(&schemas.Member{}).Where("github_id = ?", member.GithubID).Updates(map[string]any{"url": member.URL, "login": member.Login, "team_id": member.TeamID}).RowsAffected == 0 {
		db.Create(&member)
	}
}

func getDbMembers() {
	db := config.GetDatabase()
	var members []schemas.Member
	db.Preload("Team").Find(&members)
	for _, member := range members {
		fmt.Printf("List of users: %s. with (URL, TeamName, GithubID)(%s, %s, %s)\n", member.Login, member.URL, member.Team.Name, member.GithubID)
	}
}

func getTeams() ([]schemas.Team, error) {
	db := config.GetDatabase()
	var teams []schemas.Team
	result := db.Find(&teams)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		teamNameList := strings.Split(os.Getenv("FORJA_TEAMS"), ",")
		var newTeams []schemas.Team

		for _, teamName := range teamNameList {
			newTeams = append(newTeams, schemas.Team{Name: teamName})
			fmt.Printf("Creating team: %s\n", teamName)
		}

		db.Create(&newTeams)

		return getTeams()
	}

	return teams, nil
}
