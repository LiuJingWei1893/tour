package cmd

import (
	"log"

	"github.com/go-programming-tour-book/tour/internal/timer"
	"github.com/spf13/cobra"
)

// var calculateTime string
// var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNewTime()
		log.Printf("当前时间为：%s, Unix格式时间为：%d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}
