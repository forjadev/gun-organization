package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type GithubAPIRequest struct {
}

type GithubAPIResponse struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Sha         string `json:"sha"`
	Size        int    `json:"size"`
	URL         string `json:"url"`
	DownloadUrl string `json:"download_url"`
	Type        string `json:"type"`
	Content     string `json:"content"`
	Encoding    string `json:"encoding"`
}

func GetReadmeFromGithubHandler(c *gin.Context) {
	req, err := http.NewRequest("GET", "https://api.github.com/repos/Bran00/media-simples/contents/README.md", nil)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	//req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	client := http.DefaultClient
	response, err := client.Do(req)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	var apiResponse GithubAPIResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, apiResponse)
}
