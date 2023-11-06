/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// flagExampleCmd represents the flagExample command

var author string
var flagExampleCmd = &cobra.Command{
	Use:   "flagExample",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("flagExample called")
		fmt.Println("author : ", author)
	},
	//TraverseChildren: true,
}

func init() {
	rootCmd.AddCommand(flagExampleCmd)

	// Here you will define your flags and configuration settings.
	rootCmd.PersistentFlags().StringVar(&author, "author", "YOUR NAME", "Author name for copyright attribution")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// flagExampleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// flagExampleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
