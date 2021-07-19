package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"woodpecker/logger"
)

func init() {
	rootCmd.AddCommand(
		cmdResource,
		cmdStart,
	)
}

var rootCmd = &cobra.Command{
	Use:   "woodpecker",
	Short: "woodpecker for health check you know.",
	Long:  `woodpecker for health check you know.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info(fmt.Sprintf("%s", strings.Join(args, " ")))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
