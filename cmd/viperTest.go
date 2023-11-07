/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var (
	newHost string
	newPort int
)

// viperTestCmd represents the viperTest command
var viperTestCmd = &cobra.Command{
	Use:   "viperTest",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("viperTest called")

		fmt.Println("newHost : ", newHost)
		fmt.Println("newPort : ", newPort)

		// 특정 키에 대한 설정을 가져옴
		host := v.GetString("server.host")
		port := v.GetInt("server.port")

		fmt.Println("oldHost : ", host)
		fmt.Println("oldPort : ", port)
		v.Set("server.host", newHost) // set host
		v.Set("server.port", newPort) // set port

		v.WriteConfig()
		host = v.GetString("server.host")
		port = v.GetInt("server.port")

		fmt.Println("oldHost : ", host) // re-set host
		fmt.Println("oldPort : ", port) // re-set port

		// 설정을 사용하여 어떤 작업을 수행
		// 신규 설정으로 노출 되지만, 세팅한 config 파일의 수정이 없다(= viper.Set을 이용해 value를 수정할 필요가 있는 경우, file에 대해 직접 수정을 진행해야 한다.)
		// v.WriteConfig()을 사용하지 않으면, 메모리 상의 config에 대해서만 변경됨
		fmt.Printf("Server will start at %s:%d\n", host, port)

	},
}

func init() {
	rootCmd.AddCommand(viperTestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viperTestCmd.PersistentFlags().String("foo", "", "A help for foo")

	viperTestCmd.PersistentFlags().StringVar(&newHost, "newHost", "localhost", "Config Host Setting")
	viper.BindPFlag("newHost", viperTestCmd.PersistentFlags().Lookup("newHost"))

	viperTestCmd.PersistentFlags().IntVar(&newPort, "newPort", 8080, "Config Port Setting")
	viper.BindPFlag("newPort", viperTestCmd.PersistentFlags().Lookup("newPort"))

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viperTestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
