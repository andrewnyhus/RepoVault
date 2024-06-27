package github

import (
	// "fmt"
)

func ListStarredRepos() {
	PerformGETRequest("user/starred")
	// starred_repos_data := PerformGETRequest("user/starred")
	// if starred_repos_data != nil {
	// 	fmt.Printf("starred repos: %s", starred_repos_data)
	// } else {
	// 	fmt.Println("No starred repo data")
	// }
}

func ListOwnedRepos() {
}

