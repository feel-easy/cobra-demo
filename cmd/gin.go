/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"cobra-demo/gin"

	"github.com/spf13/cobra"
)

// ginCmd represents the gin command
var ginCmd = &cobra.Command{
	Use:   "gin",
	Short: "",
	Long:  ``,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("gin called")
	// },
}

var runHelloCmd = &cobra.Command{
	Use:   "hello",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		gin.RunHello()
	},
}

func init() {

	ginCmd.AddCommand(runHelloCmd)
	rootCmd.AddCommand(ginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
