Idea:

Create a Go cli that goes through directories, finds git repositories, and `git fetch`s them.
I'm already manually doing this in a bunch of repositories, automate that would be nice, and I can even learn some Go in the process.

Process:

- Make it work with a fixed directory.
- Make a cli out of it.

Commands:

- `git-sync`: syncs the current directory git repositories (1st level folders only).
- `git-sync <path>`: adds support to any (?) glob expression.
- `--recursive`: goes deeper than just 1 level in the folders tree.
- `--pull`: also pulls the current branch.

Tasks:

- [ ] List directories based on the path informed.
- [ ] Figures it out if it's a git repository.
- [ ] Executes a shell command on that directory.

Features:

- Customizing the git command.
- Block listing folders.
- Selecting only speficic folders.
- Read a file with repository paths on it, and sync only those.
