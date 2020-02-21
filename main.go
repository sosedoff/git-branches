package main

import (
	"fmt"
	"io"
	"os"

	"github.com/olekukonko/tablewriter"
)

func fatal(err interface{}) {
	fmt.Println("error:", err)
	os.Exit(1)
}

func renderBranches(branches []branchInfo, writer io.Writer) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Behind", "Ahead", "Last Commit", "Status"})

	for _, branch := range branches {
		table.Append(branch.strings())
	}

	table.Render()
}

func main() {
	branches, err := getBranches()
	if err != nil {
		fatal(err)
	}
	renderBranches(branches, os.Stdout)
}
