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

	repos := []string{}
	repos = findRepos(absPath, repos)

	if len(repos) == 0 {
		fmt.Println("no git repositories found on", absPath)
	} else {
		fmt.Println("updating projects on", absPath)
	}

	for i, repo := range repos {
		fmt.Printf("%d of %d %s\n", i+1, len(repos), repo)

		out, err := exec.Command("git", "-C", repo, "fetch", "-p", "-a").CombinedOutput()
		fmt.Printf("%v\n", string(out))
		if err != nil {
			fmt.Printf("failed to fetch: %v\n\n", err)
			continue
		}

		out, err = exec.Command("git", "-C", repo, "pull").CombinedOutput()
		fmt.Printf("%v\n", string(out))
		if err != nil {
			fmt.Printf("failed to pull: %v\n\n", err)
			continue
		}

		out, err = exec.Command("git", "-C", repo, "submodule", "update", "--recursive").CombinedOutput()
		fmt.Printf("%v\n", string(out))
		if err != nil {
			fmt.Printf("failed to update submodule: %v\n\n", err)
			continue
		}
	}

	fmt.Println()
}

func findRepos(path string, repos []string) []string {
	if isRepo(path) {
		return []string{path}
	}

	return findReposOnPath(path, repos)
}

func findReposOnPath(path string, repos []string) []string {
	dirs, _ := os.ReadDir(path)

	for _, dir := range dirs {
		if dir.IsDir() {
			fullPath := filepath.Join(path, dir.Name())

			if isRepo(fullPath) {
				repos = append(repos, fullPath)
			} else {
				repos = findReposOnPath(fullPath, repos)
			}
		}
	}

	return repos
}

func isRepo(path string) bool {
	return folderExists(filepath.Join(path, ".git"))
}

func folderExists(path string) bool {
	_, err := os.Stat(path)
	hasFolder := !os.IsNotExist(err)
	return hasFolder
}
