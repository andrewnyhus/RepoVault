package main

import (
	"fmt"
	"github.com/andrewnyhus/RepoVault/github"
	// "time"
)

func main() {
	fmt.Println("HI")
	github.ListStarredRepos()
	fmt.Println("BYE")
	// TODO start server
	// for true {
	// 	fmt.Println("HI")
	// 	time.Sleep(3 * time.Second)
	// }
}
