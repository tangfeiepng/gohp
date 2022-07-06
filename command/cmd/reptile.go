package cmd

import (
	"Walker/app/services"
	"github.com/spf13/cobra"
)

var reptileCmd = &cobra.Command{
	Use:   "reptile",
	Short: "一个简单的爬虫",
	Long:  "一个简单的针对www.zmccx.com的爬虫",
	Run: func(cmd *cobra.Command, args []string) {
		//启动爬虫
		reptile := services.Reptile{}
		reptile.Start()
	},
}
