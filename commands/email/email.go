package email

import "github.com/spf13/cobra"

func Commands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "email",
		Short: "email from the command line",
	}

	cmd.AddCommand(sendCommand())

	cmd.PersistentFlags().String("from", "", "email address to send from")
	cmd.PersistentFlags().String("to", "", "email address to send to")
	cmd.PersistentFlags().String("subject", "", "email subject")
	cmd.PersistentFlags().String("smtp", "localhost:25", "smtp server address")

	return cmd
}
