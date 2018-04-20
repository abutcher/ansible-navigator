package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/csrwng/ansible-navigator/pkg/navigator"
)

func NewAnsibleNavigatorCmd() *cobra.Command {
	return &cobra.Command{
		Use: "ansible-navigator FILENAME ROW COLUMN",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 3 {
				cmd.Usage()
				return
			}
			err := navigate(args[0], args[1], args[2], os.Stdout)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
		},
	}
}

func navigate(fileName, rowStr, colStr string, out io.Writer) error {
	row, err := strconv.Atoi(rowStr)
	if err != nil {
		return err
	}

	col, err := strconv.Atoi(colStr)
	if err != nil {
		return err
	}

	nav := &navigator.AnsibleNavigator{
		File:   fileName,
		Row:    row,
		Column: col,
	}
	path, err := nav.Navigate()
	if err != nil {
		return err
	}
	if len(path) > 0 {
		fmt.Fprintf(out, "%s", path)
	}
	return nil
}

func main() {
	cmd := NewAnsibleNavigatorCmd()
	cmd.Execute()
}
