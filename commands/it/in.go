package it

import (
	"fmt"
	"log/slog"
	"slices"

	"github.com/erikgeiser/promptkit/selection"
	"github.com/spf13/cobra"
)

func inCommand() *cobra.Command {
	inChoices := []string{"CLI", "Website", "Microservice", "Library", "Other"}

	cmd := &cobra.Command{
		Use:   "in",
		Short: "Getting user input",
		Long:  "This command demonstrates getting user input using bubbletea.",
		RunE: func(cmd *cobra.Command, args []string) error {
			choice, _ := cmd.Flags().GetString("choice")

			if choice == "" || !slices.Contains(inChoices, choice) {
				sp := selection.New("What are you building?", inChoices)
				selectedChoice, err := sp.RunPrompt()
				if err != nil {
					return fmt.Errorf("failed to get user choice: %w", err)
				}
				choice = selectedChoice
			}

			slog.Info("Input selection testing", "choice", choice)

			return nil
		},
	}

	cmd.Flags().String("choice", "", fmt.Sprintf("Choose one of the following: %v", inChoices))

	return cmd
}
