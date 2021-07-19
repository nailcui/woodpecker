package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"woodpecker/logger"
)

var isUp = false
var isLower = false
var separator = " "

func init() {
	cmdResource.PersistentFlags().BoolVarP(&isUp, "up", "u", false, "转大写输出")
	cmdResource.PersistentFlags().BoolVarP(&isLower, "lower", "l", false, "转小写输出")
	cmdResource.PersistentFlags().StringVarP(&separator, "separator", "s", " ", "自定义分隔符")
	cmdResource.AddCommand(cmdResourceReload)
}

var cmdResource = &cobra.Command{
	Use:   "resource",
	Short: "resource operation",
	Long: `resource operation, example:
		woodpecker resource reload: reload all resource`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info(fmt.Sprintf("cmd: resource, args: %v\n", args))
	},
}

var cmdResourceReload = &cobra.Command{
	Use:   "reload",
	Short: "reload resource",
	Long:  "reload resource",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info(fmt.Sprintf("reload resource success. args: %v\n", args))
	},
}
