package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println(time.Now())

	root := "/Users/muriel/projects"
	repos := []string{}
	repos = findReposOnPath(root, repos)
	fmt.Println(len(repos))

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
				log("%s\n", fullPath)
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

func log(format string, a ...any) (n int, err error) {
	debug := true
	if debug {
		return fmt.Fprintf(os.Stdout, format, a...)
	}
	return
}
