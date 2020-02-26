package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

var (
	statusActive = "ACTIVE"
	statusStale  = "STALE"
	statusDead   = "DEAD"

	colorActive = color.GreenString
	colorStale  = color.YellowString
	colorDead   = color.RedString

	staleCommitsThreshold = 100
	staleDaysThreshold    = 14

	deadCommitsThreshold = 500
	deadDaysThreshold    = 90
)

type branchInfo struct {
	name          string
	lastCommit    time.Time
	commitsAhead  int
	commitsBehind int
}

func (info branchInfo) strings() []string {
	return []string{
		info.name,
		fmt.Sprintf("%d", info.commitsBehind),
		fmt.Sprintf("%d", info.commitsAhead),
		info.lastCommit.Format(time.RFC822),
		info.status(),
	}
}

func (info branchInfo) daysSinceLastCommit() int {
	return int(time.Now().Sub(info.lastCommit).Hours()) / 24
}

func (info branchInfo) isDead() bool {
	return info.commitsBehind >= deadCommitsThreshold || info.daysSinceLastCommit() >= deadDaysThreshold
}

func (info branchInfo) isStale() bool {
	return info.commitsBehind >= staleCommitsThreshold || info.daysSinceLastCommit() >= staleDaysThreshold
}

func (info branchInfo) status() string {
	if info.isDead() {
		return colorDead(statusDead)
	}
	if info.isStale() {
		return colorStale(statusStale)
	}
	return colorActive(statusActive)
}
