# Git Sync

## Idea

Create a CLI that finds git repositories in a path and `git fetch`s and `git pull`s them.

I'm already manually doing this in a bunch of repositories, automating that would be nice, and I can even learn some Go in the process.

### Commands

- `git-sync`: syncs the current directory git repositories (all folders levels).
- `git-sync <path>`: adds support to any (?) glob expression.

### Tasks

- [x] List directories based on a fixed path.
- [x] Figures it out if it's a git repository.
- [x] Collect all repositories to be synced.
- [x] Execute a shell command on that directory.
- [x] Sync one directory at a time.
- [ ] Add path support.
- [ ] Don't try to pull when the repo is not in a clean state.
- [ ] Create a CLI.
- [ ] Release it with https://goreleaser.com/.
- [ ] Use some Go concurrecy shenanigans to speed up the process.
- [ ] Refine API.
- [ ] Use git remote instead of .git folder.
- [ ] Add flag for pulling or not?
- [ ] Add flat for remote?

### Features

- Customizing the git command.
- Block listing folders.
- Selecting only speficic folders.
- Read a file with repository paths on it, and sync only those.

### Development

Helper command to keep running the program as you make changes to it.

```bash
ls main.go | entr go run main.go
```

> https://github.com/eradman/entr
