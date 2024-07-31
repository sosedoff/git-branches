package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func main() {
	if !repositoryDetected() {
		fatal("must be in a git repo")
	}

	var filterFn filterFunc

	args := os.Args[1:]
	if len(args) > 0 {
		filterFn = func(input string) bool {
			return strings.Contains(input, args[0])
		}
	}

	branches, err := getBranches(filterFn)
	if err != nil {
		fatal(err)
	}

	renderBranches(branches, os.Stdout)
}

func fatal(err interface{}) {
	fmt.Println("error:", err)
	os.Exit(1)
}

func renderBranches(branches []branchInfo, writer io.Writer) {
	if len(branches) == 0 {
		fmt.Fprintln(writer, "No branches found")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Behind", "Ahead", "Last Commit", "Status"})

	for _, branch := range branches {
		table.Append(branch.strings())
	}

	table.Render()
}
