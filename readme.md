# Git Sync

## Idea

Create a CLI that finds git repositories in a path and `git fetch`s them.

I'm already manually doing this in a bunch of repositories, automate that would be nice, and I can even learn some Go in the process.

### Commands

- `git-sync`: syncs the current directory git repositories (1st level folders only).
- `git-sync <path>`: adds support to any (?) glob expression.
- `--recursive`: goes deeper than just 1 level in the folders tree.
- `--pull`: also pulls the current branch.

### Tasks

- [x] List directories based on a fixed path.
- [x] Figures it out if it's a git repository.
- [x] Collect all repositories to be synced.
- [ ] Execute a shell command on that directory.
- [ ] Sync one directory at a time.
- [ ] Use some Go concurrecy shenanigans to speed up the process.
- [ ] Create a CLI.
- [ ] Release it with https://goreleaser.com/.
- [ ] Refine API ideas.
- [ ] Implement API ideas.

### Features

- Customizing the git command.
- Block listing folders.
- Selecting only speficic folders.
- Read a file with repository paths on it, and sync only those.

### Development

Helper command to keep running the program as a make changes to it.

```bash
ls main.go | entr go run main.go
```

> https://github.com/eradman/entr
