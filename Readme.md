# gitstats cli tool

Gitstats is a command line tool that generates statistics from git repositories. It is written in Go and is available for Linux, Windows and Mac OS X.
The tool shows the number of commits to your local git repository similar to how github shows commit stats on your profile page.

## Installation

1. Download the repo
2. Naviate to the repo directory
3. Run the following command

```bash
    go install
```

## Usage

```bash
    gitstats path/to/repo -email=your@email.com
```

- the path to the repo is not required, it defaults to the current path
- to set a default email in the config file **~/.gitstats** add the following line

```bash
    email="your@mail.com"
```
