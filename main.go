package main

import (
	"log/slog"
	"os"
	"runtime/debug"

	"github.com/renevo/rub/commands/email"
	"github.com/renevo/rub/commands/it"

	"github.com/lmittmann/tint"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
)

func main() {
	if err := mainErr(); err != nil {
		slog.Error("rub error", "error", err)
		os.Exit(1)
	}
}

func mainErr() error {
	bi, _ := debug.ReadBuildInfo()
	rootCommand := &cobra.Command{
		Use:     "rub",
		Short:   "rub is a collection of tools by RenEvo",
		Version: bi.Main.Version,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			debug, _ := cmd.Flags().GetBool("debug")
			json, _ := cmd.Flags().GetBool("json")
			noColor, _ := cmd.Flags().GetBool("no-color")

			// logger setup
			var logLeveler slog.LevelVar
			var logHandler slog.Handler
			logOutput := os.Stderr

			switch {
			case json:
				logHandler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: &logLeveler})
			case noColor:
				logHandler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: &logLeveler})
			default:
				logHandler = tint.NewHandler(colorable.NewColorable(logOutput), &tint.Options{
					Level:   &logLeveler,
					NoColor: !isatty.IsTerminal(logOutput.Fd()),
				})
			}

			if debug {
				logLeveler.Set(slog.LevelDebug)
			}

			slog.SetDefault(slog.New(logHandler))

		},

		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	rootCommand.PersistentFlags().BoolP("debug", "d", false, "Enable application debug logging output")
	rootCommand.PersistentFlags().BoolP("json", "j", false, "Enable JSON logging output")
	rootCommand.PersistentFlags().Bool("no-color", false, "Disable colorized output on text")

	rootCommand.AddCommand(email.Commands())
	rootCommand.AddCommand(it.Commands())

	return rootCommand.Execute()
}
