/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ssargent/lawnchair/pkg/config"
)

// yamlCmd represents the yaml command
var yamlCmd2 = &cobra.Command{
	Use:   "yaml",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("yaml called")
	},
}

type yamlCmd struct {
	config     *config.Config
	configPath string
}

func init() {
	y := &yamlCmd{}
	c := &cobra.Command{
		Use:   "yaml",
		Short: "Convert config to yaml",
		Long:  `Convert config to yaml`,
		RunE:  y.run,
	}

	c.Flags().StringVarP(&y.configPath, "config", "c", "./config.json", "Path to config file")

	rootCmd.AddCommand(c)
}

func (y *yamlCmd) run(cmd *cobra.Command, args []string) error {
	cfg, err := config.New(y.configPath)
	if err != nil {
		return fmt.Errorf("yaml: %w", err)
	}

	y.config = cfg
	out, err := y.config.AsYaml()
	if err != nil {
		return fmt.Errorf("yaml: %w", err)
	}

	fmt.Println(out)
	return nil
}
