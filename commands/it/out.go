package it

import (
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"
)

func outCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "out",
		Short: "Output of data",
		Long:  "Demonstrates output of data",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := slog.Default()

			for i, arg := range args {
				logger = logger.With(fmt.Sprintf("%d", i), arg)
			}

			logger.Info("heh")

			return nil
		},
	}

	return cmd
}
