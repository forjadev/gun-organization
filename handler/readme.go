package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Member struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

const (
	orgName   = "ForjaDev"
	repoOwner = "Larissa Campos"
	repoName  = "LPMLarica"
	token     = "your_github_access_token"
)

var orgMembers = []Member{
	{Name: "John Doe", Role: "Software Engineer"},
	{Name: "Jane Smith", Role: "DevOps Engineer"},
	// Add more members as needed
}

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

func GetOrgMembersHandler(c *gin.Context) {
	// //w.Header().Set("Content-Type", "application/json")
	// err := json.NewEncoder().Encode(orgMembers)
	// if err != nil {
	// 	return
	// }
}

func UpdateReadmeHandler(c *gin.Context) {
	// members, err := getOrgMembersFromAPI()
	// if err != nil {
	// 	http.Error(w, "Failed to fetch organization members", http.StatusInternalServerError)
	// 	return
	// }

	// updatedContent := generateReadmeContent(members)

	// err = updateReadmeFile(updatedContent)
	// if err != nil {
	// 	http.Error(w, "Failed to update README.md file", http.StatusInternalServerError)
	// 	return
	// }

	// fmt.Fprint(w, "README.md file updated successfully.")
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

func getOrgMembersFromAPI() ([]Member, error) {
	resp, err := http.Get("http://localhost:8080/org-members")
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Failed to close the response body:", err)
		}
	}(resp.Body)

	var members []Member
	err = json.NewDecoder(resp.Body).Decode(&members)
	if err != nil {
		return nil, err
	}

	return members, nil
}

func generateReadmeContent(members []Member) string {
	template := `
# Organization Members

{{range .}}
- {{.Name}} ({{.Role}})
{{end}}
`

	builder := strings.Builder{}
	tmpl := strings.NewReplacer("{{range .}}", "{{range .members}}", "{{.Name}}", "%s", "{{.Role}}", "%s")
	builder.WriteString(fmt.Sprintf(template, ""))
	for _, member := range members {
		builder.WriteString(fmt.Sprintf(tmpl.Replace(template), member.Name, member.Role))
	}

	return builder.String()
}

func updateReadmeFile(updatedContent string) error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	fileContent, _, _, err := client.Repositories.GetContents(ctx, repoOwner, repoName, "README.md", nil)
	if err != nil {
		return fmt.Errorf("failed to retrieve the existing README.md file: %v", err)
	}

	content, err := fileContent.GetContent()
	if strings.TrimSpace(content) == strings.TrimSpace(updatedContent) {
		log.Println("No changes in the README.md file.")
		return nil
	}

	file, err := ioutil.TempFile("", "updated-readme-*.md")
	if err != nil {
		return fmt.Errorf("failed to create a temporary file: %v", err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {

		}
	}(file.Name())

	if _, err := file.WriteString(updatedContent); err != nil {
		return fmt.Errorf("failed to write the updated content to the temporary file: %v", err)
	}

	err = file.Close()
	if err != nil {
		return err
	}

	cmd := exec.Command("git", "checkout", "main")
	cmd.Dir = "./" // Change to your repository directory
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute 'git checkout main': %v", err)
	}

	cmd = exec.Command("git", "add", "README.md")
	cmd.Dir = "./" // Change to your repository directory
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute 'git add README.md': %v", err)
	}

	cmd = exec.Command("git", "commit", "-m", "Update README.md with organization members")
	cmd.Dir = "./" // Change to your repository directory
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute 'git commit -m': %v", err)
	}

	cmd = exec.Command("git", "push", "origin", "main")
	cmd.Dir = "./" // Change to your repository directory
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute 'git push origin main': %v", err)
	}

	return nil
}
