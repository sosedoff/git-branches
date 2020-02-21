package main

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"time"
)

func getHead() (string, error) {
	out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func parseBranchInfo(line string) (*branchInfo, error) {
	var name, date string

	if _, err := fmt.Sscanf(line, "branch=%q date=%q", &name, &date); err != nil {
		return nil, err
	}

	lastCommit, err := time.Parse("Mon Jan _2  15:04:05 2006 -0700", date)
	if err != nil {
		return nil, err
	}

	info := &branchInfo{
		name:       name,
		lastCommit: lastCommit,
	}

	behindOut, err := exec.Command("git", "rev-list", info.name+"..master").Output()
	if err != nil {
		return nil, err
	}

	aheadOut, err := exec.Command("git", "rev-list", "master.."+info.name).Output()
	if err != nil {
		return nil, err
	}

	info.commitsBehind = len(strings.Split(strings.TrimSpace(string(behindOut)), "\n"))
	info.commitsAhead = len(strings.Split(strings.TrimSpace(string(aheadOut)), "\n"))

	return info, nil
}

func getBranches() ([]branchInfo, error) {
	args := []string{
		"branch",
		"--list",
		"--no-color",
		"--sort=-committerdate",
		`--format=branch="%(refname:lstrip=2)" date="%(committerdate)"`,
	}

	out, err := exec.Command("git", args...).Output()
	if err != nil {
		return nil, err
	}

	str := strings.TrimSpace(string(out))
	lines := strings.Split(str, "\n")

	wg := sync.WaitGroup{}
	wg.Add(len(lines))

	branches := make([]branchInfo, len(lines))

	for i, line := range lines {
		go func(line string, i int) {
			defer wg.Done()

			info, err := parseBranchInfo(line)
			if err != nil {
				fatal(err)
			}
			branches[i] = *info

		}(line, i)
	}

	wg.Wait()

	return branches, nil
}
