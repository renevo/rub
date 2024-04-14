package it

import "github.com/spf13/cobra"

func Commands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "it",
		Short: "Interactive terminal command testing",
	}

	cmd.AddCommand(inCommand())
	cmd.AddCommand(outCommand())

	return cmd
}
