/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

var (
	author string
	result ExecuteResult
	v      *viper.Viper
)

// ExecuteResult는 커맨드 실행 결과를 나타내는 인터페이스입니다.
type ExecuteResult interface {
	Serialize() ([]byte, error)
}

type EmptyExecuteResult struct {
}

func (r *EmptyExecuteResult) Serialize() ([]byte, error) {
	return json.Marshal(r)
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
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 이 함수는 모든 subcommand 실행 전에 호출됩니다.
		fmt.Println("Before every subcommand...")
		result = &EmptyExecuteResult{}

	},
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
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// 이 함수는 모든 subcommand 실행 후에 호출됩니다.
		fmt.Println("After every subcommand...")
		if result != nil {
			serializedResult, err := result.Serialize()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(string(serializedResult))
			}
		}

		serializedResult2, err := json.Marshal(result)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(string(serializedResult2))
		}

		serializedResult3, err := json.Marshal(nil)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(string(serializedResult3))
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	defer func() {
		r := recover()
		fmt.Println(r)
		fmt.Println("panic recovered")
	}()
	err := rootCmd.Execute()
	fmt.Println("Root Cmd Execute end!")
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	fmt.Println("root.go init called")
	// Viper 인스턴스 설정
	v = viper.New()

	// 설정 파일 이름과 위치를 설정
	v.SetConfigName("config") // 설정 파일의 이름 (확장자 제외)
	v.SetConfigType("yaml")   // 예를 들어 "yaml" 설정 파일 형식
	v.AddConfigPath("/Users/nhn")

	// 환경 변수를 자동으로 읽어들이게 함
	v.AutomaticEnv()

	// 설정 파일을 읽어들임
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// 특정 키에 대한 설정을 가져옴
	host := v.GetString("server.host")
	port := v.GetInt("server.port")

	// 설정을 사용하여 어떤 작업을 수행
	fmt.Printf("Server will start at %s:%d\n", host, port)

	fmt.Println("root.go init ended")
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
