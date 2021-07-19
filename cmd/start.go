package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"path/filepath"
	"woodpecker/logger"
	"woodpecker/server"
	"woodpecker/single"
)

var configFileName string
var resourceDirName string

func init() {
	cmdStart.PersistentFlags().StringVarP(&configFileName, "config", "c", "config.yaml", "配置文件")
	cmdStart.PersistentFlags().StringVarP(&resourceDirName, "resource", "r", "resource.d", "资源路径")
}

var cmdStart = &cobra.Command{
	Use:   "start",
	Short: "start woodpecker for health check",
	Long:  `start woodpecker for health check`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info(fmt.Sprintf("cmd: start, args: %v\n", args))
		absConfigFileName, err := filepath.Abs(configFileName)
		if err != nil {
			logger.Info(fmt.Sprintf("configFileName: %s\n", configFileName))
			panic(err)
		}
		absResourceDirName, err := filepath.Abs(resourceDirName)
		if err != nil {
			logger.Info(fmt.Sprintf("resourceDirName: %s\n", resourceDirName))
			panic(err)
		}
		server.Start(&single.Config{
			ConfigFileName:  absConfigFileName,
			ResourceDirName: absResourceDirName,
		})
	},
}
