package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	fmt.Println(time.Now())

	root := "/Users/muriel/projects"
	repos := []string{}
	repos = findReposOnPath(root, repos)

	for i, repo := range repos {
		fmt.Printf("pulling %d of %d %s\n", i+1, len(repos), repo)
		_, err := exec.Command("git", "-C", repo, "pull", "-p").Output()
		if err != nil {
			fmt.Printf("failed to pull %v => %v\n", repo, err)
		}
	}

	fmt.Println()
}

func findReposOnPath(path string, repos []string) []string {
	dirs, _ := os.ReadDir(path)

	for _, dir := range dirs {
		if dir.IsDir() {
			fullPath := path + "/" + dir.Name()
			isRepo := folderExists(fullPath + "/.git")

			if isRepo {
				repos = append(repos, fullPath)
			} else {
				repos = findReposOnPath(fullPath, repos)
			}
		}
	}

	return repos
}

func folderExists(path string) bool {
	_, err := os.Stat(path)
	hasFolder := !os.IsNotExist(err)
	return hasFolder
}
