/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/ssargent/lawnchair/pkg/config"
)

type runCmd struct {
	config     *config.Config
	debug      bool
	configPath string
}

func (r *runCmd) run(cmd *cobra.Command, args []string) error {
	cfg, err := config.New(r.configPath)
	if err != nil {
		return fmt.Errorf("run: %w", err)
	}

	r.config = cfg

	programMap := make(map[int]config.Program)
	i := 1
	for n, p := range cfg.Programs {
		fmt.Printf("(%d)\t%s\t%s\n", i, n, p.Description)
		programMap[i] = p
		i++
	}

	fmt.Printf("\nEnter the number of the program to run: ")
	var n int
	fmt.Scanf("%d", &n)
	program := programMap[n]
	fmt.Printf("Running %s...\n", program.Description)

	cliArgs := strings.Join(program.Args, " ")

	appToRunCmd := exec.Command(program.Path, cliArgs)
	appToRunCmd.Env = os.Environ()

	for k, v := range program.EnvVars {
		appToRunCmd.Env = append(appToRunCmd.Env, fmt.Sprintf("%s=%s", k, v))
	}

	r.log(fmt.Sprintf("Env Vars set - running `%s %s`", program.Path, cliArgs))
	var outb, errb bytes.Buffer
	appToRunCmd.Stdout = &outb
	appToRunCmd.Stderr = &errb

	if err := appToRunCmd.Run(); err != nil {
		return fmt.Errorf("start: %w", err)
	}

	fmt.Println(outb.String(), "err:\n", errb.String())

	return nil
}

func (r *runCmd) log(msg string) {
	if r.debug {
		fmt.Println(msg)
	}
}

func init() { //nolint:gochecknoinits // required by cobra
	r := &runCmd{}

	c := &cobra.Command{
		Use:   "run",
		Short: "Run a program",
		Long:  `Run a program`,
		RunE:  r.run,
	}

	c.Flags().StringVarP(&r.configPath, "config", "c", "./config.json", "Path to config file")
	c.Flags().BoolVarP(&r.debug, "debug", "d", false, "Debug mode")

	rootCmd.AddCommand(c)
}
