package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	path, _ := os.Getwd()
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	absPath, _ := filepath.Abs(path)

	fmt.Println("updating projects on", absPath)

	repos := []string{}
	repos = findReposOnPath(absPath, repos)

	for i, repo := range repos {
		fmt.Printf("%d of %d %s\n", i+1, len(repos), repo)

		_, err := exec.Command("git", "-C", repo, "fetch", "-p", "-a").Output()
		if err != nil {
			fmt.Printf("\tfailed to fetch: %v\n", err)
			continue
		}

		_, err = exec.Command("git", "-C", repo, "pull").Output()
		if err != nil {
			fmt.Printf("\tfailed to pull: %v\n", err)
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
