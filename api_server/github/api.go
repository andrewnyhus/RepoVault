package github

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)


func PerformGETRequest(route string) {
	client := http.Client{Timeout: 3 * time.Second}
	// TODO handle pagination
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/%s", route), nil)
	// log error details if request creation failed
	if err != nil {
		fmt.Printf("Error while creating request to GitHub API (route: %s): %s", route, err)
		return
	}

	req.Header.Add("Accept", "application/json")
	// add api token header
	fmt.Printf("token: %s", os.Getenv("GITHUB_API_TOKEN"))
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GITHUB_API_TOKEN")))

	// make the request
	resp, err := client.Do(req)
	// log error details if request failed
	if err != nil {
		fmt.Printf("Error while making request to GitHub API (route: %s): %s", route, err)
		return
	}

	// close the response body
	defer resp.Body.Close()
	// read response data
	body, err := ioutil.ReadAll(resp.Body)
	// log error details if parsing failed
	if err != nil {
		fmt.Printf("Error parsing response data from GitHub API (route: %s): %s", route, err)
		return
	}

	// TODO decide whether only returning the body is sufficient or if status code and other attributes might be needed
	fmt.Printf("Response body: %s", body)
}