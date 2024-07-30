package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"time"
)

func repositoryDetected() bool {
	out := bytes.NewBuffer(nil)
	cmd := exec.Command("git", "status")
	cmd.Stderr = out
	cmd.Run()

	return !strings.Contains(out.String(), "not a git repository")
}

func getHead() (string, error) {
	out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func parseBranchInfo(branch string, line string) (*branchInfo, error) {
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

	behindOut, err := exec.Command("git", "rev-list", info.name+".."+branch).Output()
	if err != nil {
		return nil, err
	}

	aheadOut, err := exec.Command("git", "rev-list", branch+".."+info.name).Output()
	if err != nil {
		return nil, err
	}

	info.commitsBehind = len(strings.Split(strings.TrimSpace(string(behindOut)), "\n"))
	info.commitsAhead = len(strings.Split(strings.TrimSpace(string(aheadOut)), "\n"))

	return info, nil
}

func getBranches() ([]branchInfo, error) {
	current, err := getHead()
	if err != nil {
		return nil, err
	}

	merged, err := getMergedBranches(current)
	if err != nil {
		return nil, err
	}

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

			if line == current {
				return
			}

			info, err := parseBranchInfo(current, line)
			if err != nil {
				fatal(err)
			}

			for _, b := range merged {
				if b == info.name {
					info.merged = true
				}
			}

			branches[i] = *info

		}(line, i)
	}

	wg.Wait()

	return branches, nil
}

func getMergedBranches(current string) ([]string, error) {
	args := []string{
		"branch",
		"--list",
		"--no-color",
		`--format="%(refname:lstrip=2)"`,
		"--merged=" + current,
	}

	out, err := exec.Command("git", args...).CombinedOutput()
	if err != nil {
		return nil, err
	}

	str := strings.TrimSpace(string(out))
	lines := strings.Split(str, "\n")

	result := []string{}
	for _, line := range lines {
		var name string
		fmt.Sscanf(line, "%q", &name)

		if name != current {
			result = append(result, name)
		}
	}

	return result, nil
}
