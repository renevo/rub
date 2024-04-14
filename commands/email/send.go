package email

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	gomail "gopkg.in/mail.v2"
)

func sendCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send",
		Short: "send an email",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			from, _ := cmd.Flags().GetString("from")
			to, _ := cmd.Flags().GetString("to")
			subject, _ := cmd.Flags().GetString("subject")
			smtpAddr, _ := cmd.Flags().GetString("smtp")

			if from == "" || to == "" || subject == "" || smtpAddr == "" {
				return cmd.Help()
			}

			mail := gomail.NewMessage()
			mail.SetHeader("From", from)
			mail.SetHeader("To", to)
			mail.SetHeader("Subject", subject)
			mail.SetBody("text/plain", args[0])

			host, portStr, _ := net.SplitHostPort(smtpAddr)
			port, _ := strconv.Atoi(portStr)

			d := &gomail.Dialer{
				Host:         host,
				Port:         port,
				Timeout:      5 * time.Second,
				RetryFailure: true,
			}

			if err := d.DialAndSend(mail); err != nil {
				return fmt.Errorf("failed to send email: %w", err)
			}

			return nil
		},
	}

	return cmd
}
