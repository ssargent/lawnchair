/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// sample1Cmd represents the sample1 command
var sample1Cmd = &cobra.Command{
	Use:   "sample1",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sample1 called")
	},
}

func init() {
	rootCmd.AddCommand(sample1Cmd)

	// Here you will define your flags and configuration settings.
	sample1Cmd.Flags().String("config", "c", "Path to config file")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sample1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sample1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
