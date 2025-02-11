/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// sample2Cmd represents the sample2 command
var sample2Cmd = &cobra.Command{
	Use:   "sample2",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sample2 called")
		worker := os.Getenv("WORKER_HOST")
		id := os.Getenv("WORKER_ID")
		port := os.Getenv("WORKER_PORT")

		fmt.Printf("WORKER_HOST: %s\nWORKER_ID: %s\nWORKER_PORT: %s\n", worker, id, port)
	},
}

func init() {
	rootCmd.AddCommand(sample2Cmd)

	// Here you will define your flags and configuration settings.

	sample2Cmd.Flags().String("config", "c", "Path to config file")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sample2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sample2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
