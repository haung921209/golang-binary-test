/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	author string
	result ExecuteResult
)

// ExecuteResult는 커맨드 실행 결과를 나타내는 인터페이스입니다.
type ExecuteResult interface {
	Serialize() ([]byte, error)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "golang-binary-test root",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files 
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Pre run")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println(args[0])
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Post run")
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 이 함수는 모든 subcommand 실행 전에 호출됩니다.
		fmt.Println("Before every subcommand...")

	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// 이 함수는 모든 subcommand 실행 후에 호출됩니다.
		fmt.Println("After every subcommand...")
		serializedResult, err := result.Serialize()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(string(serializedResult))
		}

		serializedResult2, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(string(serializedResult2))
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.golang-binary-test.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// rootCmd.PersistentFlags().StringVar(&author, "author", "YOUR NAME", "Author name for copyright attribution")

	// runtime error due to variable scope,,,,
	// var author1 string
	// rootCmd.PersistentFlags().StringVar(&author1, "author", "YOUR NAME", "Author name for copyright attribution")
}
