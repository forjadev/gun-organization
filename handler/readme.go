package handler

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type GithubAPIRequest struct {
	Message 			string `json:"message"`
	Commiter 			GithubAPICommiter `json:"commiter"`
	Content 			string `json:"content"`
	Sha 					string `json:"sha"`
}

type GithubAPICommiter struct {
	Name 					string `json:"name"`
	Email 				string `json:"email"`
}


type GithubAPIResponse struct {
	Name           string `json:"name"`
	Path           string `json:"path"`
	Sha            string `json:"sha"`
	Size           int    `json:"size"`
	URL            string `json:"url"`
	DownloadUrl    string `json:"download_url"`
	Type           string `json:"type"`
	Content        string `json:"content"`
	Encoding       string `json:"encoding"`
	ContentDecoded string
	ContentEncoded string
}

func GetReadmeFromGithubHandler(c *gin.Context) {
	req, err := http.NewRequest("GET", "https://api.github.com/repos/Bran00/media-simples/contents/README.md", nil)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")

	client := http.DefaultClient
	response, err := client.Do(req)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()

	//TODO: Check new version to replace ioutil.ReadAll
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
	contentDecoded, _ := decodeBase64(apiResponse.Content)

	newUser := `| <img src="https://avatars.githubusercontent.com/u/20804404?v=4" width="100"> | [Vinicius Rossado](https://github.com/vinirossado) | Software Engineer |`

	lastImgIndex := strings.LastIndex(contentDecoded, "| <img src=")

	beforeLastImg := contentDecoded[:lastImgIndex]
	afterLastImg := contentDecoded[lastImgIndex:]

	contentDecoded = beforeLastImg + newUser + "\n" + afterLastImg

	apiResponse.ContentDecoded = contentDecoded
	apiResponse.ContentEncoded = encodeBase64(contentDecoded)

	var reqUpdate GithubAPIRequest
		reqUpdate.Message = "Atualizando Readme Via Golang"
		reqUpdate.Commiter.Name = "Brando Rocha"
		reqUpdate.Commiter.Email = "brandorocha00@gmail.com"
		reqUpdate.Content = apiResponse.ContentEncoded
		reqUpdate.Sha = apiResponse.Sha

	updateReadme(&reqUpdate)

	c.JSON(http.StatusOK, apiResponse)
}

func encodeBase64(data string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(data))
	return encoded
}

func updateReadme( body *GithubAPIRequest ) {
	jsonData, err := json.Marshal(body)
	req, _ := http.NewRequest("PUT", "https://api.github.com/repos/Bran00/media-simples/contents/README.md", bytes.NewBuffer(jsonData))

	/*if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	} */

	token := "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"
	req.Header.Set("Authorization", "Bearer "+ token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	response, err := client.Do(req)

	/*if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}*/

	defer response.Body.Close()

}

func decodeBase64(data string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		fmt.Println("Error decoding string: ", err)
		return "", err
	}

	decoded := string(decodedBytes)
	return decoded, nil
}
