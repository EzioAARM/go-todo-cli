/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/ezioaarm/go-todo-cli/cmd/add"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "todo",
	Short: "CLI to manage tasks in real time",
	Long: `Modern CLI to manage tasks in real time.
	
This CLI has support for multiple database engines (PostgreSQL, CockroachDB, Neo4j).`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(add.AddCmd)
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
