/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

type Example1Result struct {
	Message string `json:"message"`
}

func (r *Example1Result) Serialize() ([]byte, error) {
	return json.Marshal(r)
}

// commandExampleCmd represents the commandExample command
var commandExampleCmd = &cobra.Command{
	Use:   "COMMAND_EXAMPLE",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("commandExample called")
		fmt.Println("commandExample run started")
		result = &Example1Result{
			Message: "Example1 executed",
		}
		serializedResult, _ := result.Serialize()
		cmd.Println(string(serializedResult))
		fmt.Println("commandExample end")
	},
}

func init() {
	rootCmd.AddCommand(commandExampleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commandExampleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// commandExampleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
